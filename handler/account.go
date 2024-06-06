package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
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

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	if name != "" {
		q = q.Where("name = ?", name)
	}

	result := q.Find(&accounts)

	// result := orm.Find(&accounts, "name = ?", name)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		"data":    accounts,
	})
}

// type BodyPayLoadAccount struct {
// 	AccountID string
// 	Name      string
// 	Address   string
// }

func (a *accountImplement) CreateAccount(g *gin.Context) {
	BodyPayLoad := model.Account{}

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

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	BodyPayLoad := model.Account{}
	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}
	orm.First(&user, "account_id = ?", id)
	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	user.Name = BodyPayLoad.Name
	user.Username = BodyPayLoad.Username
	orm.Save(user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Account retrieved successfully",
		// "data":    map[string]string{"name": name},
	})

}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete Account successfully",
		"data":    map[string]string{"id": id},
	})

}
