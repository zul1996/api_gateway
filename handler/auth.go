package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	AuthLogin(*gin.Context)
}

type AuthImplement struct{}

func Login() AuthInterface {
	return &AuthImplement{}
}

type BodyPayLoadAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthImplement) AuthLogin(c *gin.Context) {
	var bodyPayLoad BodyPayLoadAuth

	err := c.BindJSON(&bodyPayLoad)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if bodyPayLoad.Username == "admin" && bodyPayLoad.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Account retrieved successfully",
			"data":    bodyPayLoad,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid username or password",
		})
	}
}
