package listeners

import (
	"burd/listeners/message"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session) {
	s.AddHandler(message.MessageCreate)
}
