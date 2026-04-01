package handlers

import (
	"example/spark/internals"
	"example/spark/internals/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var users models.User
	err := c.ShouldBindJSON(&users)
	if err != nil {
		fmt.Println("error binding user")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	internals.Users = append(internals.Users, users)
	c.JSON(http.StatusOK, map[string]any{
		"message": "user created successfully",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, internals.Users)
}

func getuser(id int) (models.User, bool) {
	for _, usr := range internals.Users {
		if id == usr.ID {
			return usr, true

		}
	}
	return models.User{}, false
}

func GetOneUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "an error occured",
		})
		fmt.Println(err)
		return
	} else {
		usr, exist := getuser(id)
		if exist {
			c.JSON(http.StatusOK, usr)
		} else {
			c.JSON(200, gin.H{
				"message": "no user exist yet",
			})
			return
		}
	}
}
