package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

// Include UserResponse struct definition here or import from its file

func main() {
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve specific user "644b2ba8cc3f875be1b7658d" under the organization "64131e880654550010d30c54"
	user_org1, err := client.UserGet(ctx, "64131e880654550010d30c54", "644b2ba8cc3f875be1b7658d") // Only provide organization ID and desired user ID
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, user := range user_org1 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve all users for organization "641351fee8f281432b807a50"
	users_org2, err := client.UserGet(ctx, "641351fee8f281432b807a50")
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, user := range users_org2 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	// JSON Output
	// userBytes, err := json.MarshalIndent(users_org2, "", "  ")
	// if err != nil {
	// 	log.Println("Error marshalling user:", err)
	// } else {
	// 	fmt.Println(string(userBytes))
	// }

}