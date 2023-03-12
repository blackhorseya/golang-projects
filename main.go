package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "bot token")
	flag.Parse()
}

func main() {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("error creating discord session, ", err)
	}

	s.AddHandler(handleMessage)

	s.Identify.Intents = discordgo.IntentsGuildMessages

	err = s.Open()
	if err != nil {
		log.Fatalln("error opening connection, ", err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signals

	s.Close()
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
