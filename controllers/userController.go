package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/BinayRajbanshi/go-auth/database"
	"github.com/BinayRajbanshi/go-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func Login(c *gin.Context) {
	// get email and password from the request body
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// check if the user exist
	// db.Where("name = ?", "jinzhu").First(&user)
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password.",
		})
		return
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password.",
		})
		return
	}

	// generate token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week expiration
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot create a token",
			"stack": err,
		})
		return
	}

	// respond back

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"exp":   time.Now().Add(time.Hour * 24 * 7),
	})
}
