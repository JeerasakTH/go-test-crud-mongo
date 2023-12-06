package controller

import (
	"net/http"

	"github.com/JeerasakTH/go-test-crud/model"
	"github.com/JeerasakTH/go-test-crud/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		userID := c.Query("id")
		objId, _ := primitive.ObjectIDFromHex(userID)

		user := model.User{}
		filter := bson.M{"_id": objId}
		filterOp := bson.M{}
		if err := repository.GetOne("test", "user_statement", filter, filterOp, &user); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"Data":    user,
		})
	}
}
