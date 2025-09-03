package main

import (
	"burd/commands"
	"burd/listeners"
	"burd/logger"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "token", "", "discord token")
	flag.Parse()
}

func main() {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Log.Fatalf("error creating client: %s", err.Error())
	}
	defer client.Close()

	client.Identify.Intents = discordgo.IntentGuildMessages

	err = client.Open()
	if err != nil {
		logger.Log.Fatalf("error opening connection: %s", err.Error())
	}
	logger.Log.Info("client started")

	listeners.Register(client)
	commands.Register(client)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	commands.Unregister(client)
}
