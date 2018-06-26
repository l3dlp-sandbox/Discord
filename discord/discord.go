package discord

import (
	"../api"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
)

var ApiUrl = ""

func Run(token, apiUrl string) {
	ApiUrl = apiUrl

	// Create a new Discord session using the support token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating the discord bot session: ", err)
		return
	}

	// Register the event
	dg.AddHandler(messageCreate)

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

// Retrieves the user entry when the discord bot is mentionned and respond with the chatbot
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Discord format for bot mention
	botMention := fmt.Sprintf("<@%s>", s.State.User.ID)

	if !strings.HasPrefix(m.Content, botMention) {
		return
	}

	s.ChannelTyping(m.ChannelID)

	// Send a request to the rest API with the given sentence
	response := api.Respond(ApiUrl, strings.Replace(m.Content, botMention, "", 1), m.Author.ID)

	// Respond it with a user mention
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> %s", m.Author.ID, response))
}
