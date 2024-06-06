package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
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

// type BodyPayLoadTransaction struct {
// 	Amount      string
// 	FromAccount string
// 	ToAccount   string
// }

func (a *transactionImplement) CreateTransaction(g *gin.Context) {
	BodyPayLoad := model.Transaction{}

	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&BodyPayLoad)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this rest api for later",
		"data":    BodyPayLoad,
	})

}
