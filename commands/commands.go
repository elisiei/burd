package commands

import (
	"burd/commands/invert"
	"burd/logger"

	"github.com/bwmarrin/discordgo"
)

var registeredCommands []*discordgo.ApplicationCommand

func Register(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{&invert.Command}
	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		invert.Command.Name: invert.Handler,
	}
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			logger.Log.Fatalf("couldn't create command %s: %s", v.Name, err.Error())
		}

		logger.Log.Debugf("created slash command %s", v.Name)
		registeredCommands[i] = cmd
	}
}

func Unregister(s *discordgo.Session) {
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			logger.Log.Fatalf("couldn't delete command %s: %s", v.Name, err.Error())
		}
	}
}
