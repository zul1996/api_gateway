package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionInterface interface {
	CreateTransaction(*gin.Context)
}

type transactionImplement struct{}

func Transaction() transactionInterface {
	return &transactionImplement{}
}

type BodyPayLoadTransaction struct {
	Amount      string
	FromAccount string
	ToAccount   string
}

func (a *transactionImplement) CreateTransaction(g *gin.Context) {
	BodyPayLoad := BodyPayLoadTransaction{}

	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return

	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this rest api for later",
		"data":    BodyPayLoad,
	})

}
