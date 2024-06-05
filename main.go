package main

import (
	"api_gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/patch/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.GET("/balance", handler.NewAccount().GetAccountBalance)

	authRoute := r.Group("/auth")
	authRoute.POST("/post", handler.Login().AuthLogin)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/post", handler.Transaction().CreateTransaction)

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
