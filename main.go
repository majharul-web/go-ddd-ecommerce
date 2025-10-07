package main

import "ecommerce/cmd"

func main() {
	cmd.Serve()

	// jwt, err := util.CreateJWT("my_secret_key", util.Payload{
	// 	Sub:         1,
	// 	FirstName:   "John",
	// 	LastName:    "Doe",
	// 	Email:       "john.doe@example.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println("Error creating JWT:", err)
	// 	return
	// }
	// fmt.Println("Generated JWT:", jwt)

}
