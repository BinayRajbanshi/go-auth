package controllers

import (
	"net/http"

	"github.com/BinayRajbanshi/go-auth/database"
	"github.com/BinayRajbanshi/go-auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get credentials from body
	var body struct {
		Email    string `validate:"email"`
		Password string `validate:"min=8"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to read body",
		})
		return
	}
	// hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create a user i.e. save info in database
	user := models.User{Email: body.Email, Password: string(hash)}

	result := database.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to create a user",
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"success": "user created successfully",
	})

}
