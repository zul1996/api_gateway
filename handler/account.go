package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	GetAccountBalance(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccountBalance(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	balance := queryParam.Get("balance")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		"data":    map[string]string{"name": balance},
	})

}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		"data":    map[string]string{"name": name},
	})

}

type BodyPayLoadAccount struct {
	AccountID string
	Name      string
	Address   string
}

func (a *accountImplement) CreateAccount(g *gin.Context) {
	BodyPayLoad := BodyPayLoadAccount{}

	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this rest api for later",
		"data":    BodyPayLoad,
	})

}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		"data":    map[string]string{"name": name},
	})

}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		"data":    map[string]string{"name": id},
	})

}
