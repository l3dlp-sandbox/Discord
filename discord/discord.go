package discord

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run(token string) {
	// Create a new Discord session using the support token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating the discord bot session: ", err)
		return
	}

	// Open the connection
	err = dg.Open()
	if err != nil {
		fmt.Println("Rrror opening connection: ", err)
		return
	}

	// Wait here until CTRL-C
	fmt.Println("Discord support is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
