package main

import (
	"api_gateway/handler"
	"context"
	"net/http"

	"api_gateway/proto"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	addrServiceTransactionOpt := client.WithAddress(":9000")
	clientSrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/patch/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.Login().AuthLogin)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/create", handler.Transaction().CreateTransaction)
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).Login(context.Background(), &proto.LoginRequest{
			Username: "zoel",
			Password: "zoel123",
		}, addrServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// package main

// import (
// 	"api_gateway/usecase"
// 	"github.com/gin-gonic/gin"
// 	"fmt"
// )

// func main() {

// 	Login := usecase.NewLogin("basic")
// 	if Login != nil {
// 		fmt.Println("Login user:", Login.Autentikasi("admin", "admin123"))     // Harusnya true
// 		fmt.Println("Login user:", Login.Autentikasi("user", "wrongpassword")) // Harusnya false
// 	} else {
// 		fmt.Println("Invalid login type for basic")
// 	}

// }
