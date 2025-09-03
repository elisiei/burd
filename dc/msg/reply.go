package msg

import "github.com/bwmarrin/discordgo"

// simple reply, string only
func Reply(s *discordgo.Session, i *discordgo.InteractionCreate, c string, data *discordgo.InteractionResponseData) error {
	data.Content = c
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
}
