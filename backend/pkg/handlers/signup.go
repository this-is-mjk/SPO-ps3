package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/this-is-mjk/SPO/p3/pkg/helpers"
	"github.com/this-is-mjk/SPO/p3/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *gin.Context, users *mongo.Collection) {
	// fmt.Println("Sign up")
	var userData models.User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, "Invalid request, please fill all the fields")
		return
	}
	hashedUserId, _ := helpers.Hash(userData.UserId)
	_, exists := helpers.FindUser(users, bson.M{"userId": hashedUserId})
	if exists {
		c.JSON(http.StatusOK, "User already exists, Sign in instead")
		return
	}
	// fmt.Println(registrationData)
	hashPassword, _ := helpers.Hash(userData.Password)
	helpers.InsertUser(users, bson.M{"userId": hashedUserId, "password": hashPassword})
	c.JSON(http.StatusOK, "Registration done")
}
