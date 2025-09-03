package invert

import "github.com/bwmarrin/discordgo"

var Command = discordgo.ApplicationCommand{
	Name:        "invert",
	Description: "invert an image",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "image",
			Description: "image to process",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}
