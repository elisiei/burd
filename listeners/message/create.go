package message

import (
	"burd/logger"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		logger.Log.Debug("ignored bot command")
		return
	}
}
