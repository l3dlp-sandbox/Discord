package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"fmt"
)

var (
	Token = os.Getenv("DISCORD_TOKEN")
	ApiAddress = os.Getenv("API_ADRESS")
)

func main() {
	// Ask the token if there isn't in the environment variable
	if Token == "" {
		fmt.Print("Enter your discord token: ")
		token, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}

		Token = string(token)
	}

	if ApiAddress == "" {
		fmt.Print("Enter your api adress: ")
		var address string
		fmt.Scan(&address)

		ApiAddress = address
	}
}