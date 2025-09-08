package main

import (
	"fmt"
	"go-practice-api/utilities"
)

func main() {
	// cmd.Serve()
	jwt, err := utilities.CreateJwt("my-secret", utilities.Payload{
		Sub:         35,
		FirstName:   "Rakibul",
		LastName:    "Hasan",
		Email:       "hasan@example.com",
		IsShopOwner: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}
