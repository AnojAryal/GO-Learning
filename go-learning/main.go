package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var username, password string

	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// Validate the username and password
	if len(username) == 0 {
		fmt.Println("Error: Username cannot be empty.")
		return
	}
	if len(password) == 0 {
		fmt.Println("Error: Password cannot be empty.")
		return
	}

	// Encode the username and password in Base64
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Print the Authorization header
	fmt.Println("Authorization: Basic", encodedAuth)
}
