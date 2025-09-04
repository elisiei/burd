package grayscale

import "github.com/bwmarrin/discordgo"

var Command = discordgo.ApplicationCommand{
	Name:        "grayscale",
	Description: "grayscale an image",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "image",
			Description: "image to process",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}
