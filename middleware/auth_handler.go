package middleware

import (
	"time"

	"go_auth/config"
	"go_auth/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var loginData LoginSchema

	if err := c.ShouldBindJSON(&loginData); err != nil {
		errors := FormatValidationErrors(err)
		c.JSON(403, gin.H{"errors": errors})
		return
	}

	validationErrors := ValidateInput(loginData)
	if validationErrors != nil {
		c.JSON(403, gin.H{"errors": validationErrors})
		return
	}

	// Cari pengguna berdasarkan email
	var user model.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(config.JWTExpireDuration()).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
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
		c.JSON(500, gin.H{"error": "Failed to invalidate token"})
		return
	}

	// Clear cookie or token
	c.JSON(200, gin.H{"message": "Logout successful", "token": tokenString})
}
