package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/this-is-mjk/SPO/p3/pkg/helpers"
	"github.com/this-is-mjk/SPO/p3/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context, users *mongo.Collection) {
	var userData models.User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, "Invalid request, please fill all the fields")
		return
	}
	hashedUserId, _ := helpers.Hash(userData.UserId)
	user, exist := helpers.FindUser(users, bson.M{"userId": hashedUserId})
	if !exist {
		c.JSON(http.StatusOK, "User not found, Sign up first")
		return
	}
	if !helpers.VerifyHash(userData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, "Invalid password")
		return
	}
	c.JSON(http.StatusOK, "Login successful")
}
