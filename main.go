package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"./discord"
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
		fmt.Println("")
	}

	if ApiAddress == "" {
		fmt.Print("Enter your api url (http://xxx.xxx): ")
		var address string
		fmt.Scan(&address)

		ApiAddress = address
	}

	discord.Run(Token, ApiAddress)
}