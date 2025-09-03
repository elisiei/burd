package opt

import "github.com/bwmarrin/discordgo"

// GetOptions returns a map of options by name.
func GetOptions(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	return optionMap
}

// Typed accessor helpers:

func String(opt *discordgo.ApplicationCommandInteractionDataOption) string {
	if opt == nil {
		return ""
	}
	return opt.StringValue()
}

func Int(opt *discordgo.ApplicationCommandInteractionDataOption) int64 {
	if opt == nil {
		return 0
	}
	return opt.IntValue()
}

func Bool(opt *discordgo.ApplicationCommandInteractionDataOption) bool {
	if opt == nil {
		return false
	}
	return opt.BoolValue()
}

func User(opt *discordgo.ApplicationCommandInteractionDataOption, s *discordgo.Session) *discordgo.User {
	if opt == nil {
		return nil
	}
	return opt.UserValue(s)
}

func Channel(opt *discordgo.ApplicationCommandInteractionDataOption, s *discordgo.Session) *discordgo.Channel {
	if opt == nil {
		return nil
	}
	return opt.ChannelValue(s)
}

func Attachment(i *discordgo.InteractionCreate, opt *discordgo.ApplicationCommandInteractionDataOption) *discordgo.MessageAttachment {
	if opt == nil {
		return nil
	}

	attachmentID := opt.Value.(string)
	attachment, ok := i.ApplicationCommandData().Resolved.Attachments[attachmentID]
	if !ok {
		return nil
	}
	return attachment
}
