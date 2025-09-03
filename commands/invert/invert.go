package invert

import (
	"burd/dc/msg"
	"burd/dc/opt"
	"burd/img/effects"
	"burd/img/io"
	"bytes"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := opt.GetOptions(i)
	image := opt.String(opts["image"])

	data, err := io.FetchBytes(image)
	if err != nil {
		msg.Reply(s, i, err.Error(), nil)
		return
	}

	decoded, err := io.Decode(data)
	if err != nil {
		msg.Reply(s, i, err.Error(), nil)
		return
	}

	result := effects.Invert(decoded)
	encoded, err := io.Encode(result, io.PNGEncoder())
	if err != nil {
		msg.Reply(s, i, err.Error(), nil)
		return
	}

	msg.Reply(s, i, "done!", &discordgo.InteractionResponseData{
		Files: []*discordgo.File{
			{
				Name:        "xd.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(encoded),
			},
		},
	})
}
