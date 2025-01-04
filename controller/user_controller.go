package controller

import (
	"fmt"
	"go_auth/config"
	"go_auth/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully", "user": user})
}

func GetAllUserHandler(c *gin.Context) {
	var user []model.User
	if err := config.DB.Find(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

func GetUserHandler(c *gin.Context) {
	// Ambil ID user dari parameter URL
	id := c.Param("id")

	// Ambil current_id dan role dari context (diset oleh middleware)
	currentUserID, exists := c.Get("current_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	role, roleExists := c.Get("role")
	if !roleExists || (role != "admin" && id != fmt.Sprint(currentUserID)) {
		c.JSON(403, gin.H{"error": "Forbidden: You can only access your own data"})
		return
	}

	// Pastikan user ada
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")

	// Ambil current_id dan role dari context (diset oleh middleware)
	currentUserID, exists := c.Get("current_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	role, roleExists := c.Get("role")
	if !roleExists || (role != "admin" && id != fmt.Sprint(currentUserID)) {
		c.JSON(403, gin.H{"error": "Forbidden: You can only access your own data"})
		return
	}

	// Pastikan user ada
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Bind data JSON yang akan di-update
	var updatedData model.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update field yang diperbolehkan
	user.Name = updatedData.Name
	user.Email = updatedData.Email

	// Jika password diperbarui, hash terlebih dahulu
	if updatedData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully", "user": user})
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	// Ambil current_id dan role dari context (diset oleh middleware)
	currentUserID, exists := c.Get("current_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	role, roleExists := c.Get("role")
	if !roleExists || (role != "admin" && id != fmt.Sprint(currentUserID)) {
		c.JSON(403, gin.H{"error": "Forbidden: You can only access your own data"})
		return
	}

	// Pastikan user ada
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Hapus user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
