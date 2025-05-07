package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/initializers"
	"github.com/yoonaji/carbon_test/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (pc *UserController) UpdateUser(c *gin.Context) {
	userId := c.Param("userId")

	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Email = input.Email
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
