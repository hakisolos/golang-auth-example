package handlers

import (
	"example/spark/internals"
	"example/spark/internals/models"
	"example/spark/internals/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {

		fmt.Println("error binding user")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// first check if user exists already:
	exits := utils.UserExists(user.Email)
	if exits {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}
	hashedPassword, err := utils.Hashpassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "an error occured",
		})
		fmt.Println(err)
		return
	}
	user.Password = hashedPassword
	internals.Users = append(internals.Users, user)
	c.JSON(http.StatusOK, map[string]any{
		"message": "user created successfully",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {
	if internals.Users == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no users found",
		})
		return
	}
	c.JSON(http.StatusOK, internals.Users)
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
		usr, exist := utils.Getuser(id)
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

func DeleteUser(c *gin.Context) {
	//check if user even existed first
	email := c.Param("email")
	exits := utils.UserExists(email)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user does not exist",
		})
	} else {
		err := utils.DeleteUser(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error deleting user, try again.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "user deleted successfully",
		})
	}
}

func HandleLogin(c *gin.Context) {
	jwtsecret := os.Getenv("JWT_SECRET")
	if jwtsecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "jwt secret not configured",
		})
		return
	}
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user loginRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}
	// now to check if user even  xists
	exits := utils.UserExists(user.Email)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user does not exist",
		})
		return
	}
	//check the password that corresponds with the database
	var password string
	for _, usr := range internals.Users {
		if usr.Email == user.Email {
			password = usr.Password
		}
	}
	login := utils.ComparePassword(user.Password, password)
	if login {
		// jwt functionality
		claims := jwt.MapClaims{
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tk, err := token.SignedString([]byte(jwtsecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to generate token",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "user logged in successfully",
			"data":    tk,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect email or password",
		})
	}

}

func Me(c *gin.Context) {
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "user not found in context"})
		return
	}

	claims := userData.(jwt.MapClaims)

	email := claims["email"]

	c.JSON(200, gin.H{
		"message": "profile fetched",
		"email":   email,
	})
}
