package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go_auth/config"
	"go_auth/model"
)

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Validasi input
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cari pengguna berdasarkan email
	var user model.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(config.JWTExpireDuration()).Unix(), // Ambil durasi dari config
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret)) // Gunakan JWTSecret dari config
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Berhasil login
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func LogoutHandler(c *gin.Context) {
	// Invalidate token by setting claims exp to time.Now() - 1 (expired token immediately)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 0,
		"role":    "",
		"exp":     -1, // Token expired instantly
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate token"})
		return
	}

	// Clear cookie or token
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful", "token": tokenString})
}