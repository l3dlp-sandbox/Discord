package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"github.com/olivia-ai/Discord/discord"
	"os"
	"fmt"
	"strings"
)

var (
	Token = os.Getenv("DISCORD_TOKEN")
	ApiAddress = os.Getenv("API_ADRESS")
)

func main() {
	// Ask the token if there isn't in the environment variable
	if Token == "" {
		fmt.Print("Enter Olivia's discord token: ")
		token, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}

		Token = string(token)
		fmt.Println("")
	}

	if ApiAddress == "" {
		fmt.Print("Do you want to use a custom api (C) or the offical (O) ? ")

		for {
			var choice string
			fmt.Scan(&choice)
			choice = strings.ToLower(choice)

			if choice == "c" {
				fmt.Print("Enter the api url (http://xxx.xxx): ")
				fmt.Scan(&ApiAddress)
				break
			} else if choice == "o" {
				ApiAddress = "https://olivia-api.herokuapp.com"
				break
			} else {
				fmt.Print("Please enter C for a custom api or O for the official: ")
			}
		}
	}

	discord.Run(Token, ApiAddress)
}