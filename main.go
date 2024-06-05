package main

import (
	"api_gateway/usecase"
	"fmt"
)

func main() {

	Login := usecase.NewLogin("basic")
	if userLogin != nil {
		fmt.Println("Login user:", Login.Autentikasi("admin", "admin123"))     // Harusnya true
		fmt.Println("Login user:", Login.Autentikasi("user", "wrongpassword")) // Harusnya false
	} else {
		fmt.Println("Invalid login type for basic")
	}

}
