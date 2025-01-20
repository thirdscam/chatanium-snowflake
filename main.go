package main

import (
	"fmt"
	"math/big"
	"time"

	"antegr.al/chatanium-bot/v1/src/Backends/Discord/Interface/Slash"
	"antegr.al/chatanium-bot/v1/src/Util/Log"
	"github.com/bwmarrin/discordgo"
)

var MANIFEST_VERSION = 1

var (
	NAME       = "Snowflake"
	BACKEND    = "discord"
	VERSION    = "0.0.1"
	AUTHOR     = "ANTEGRAL"
	REPOSITORY = "github:thirdscam/chatanium"
)

var DEFINE_SLASHCMD = Slash.Commands{
	{
		Name:        "s2t",
		Description: "Convert Snowflake ID to time",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "snowflake",
				Description: "Enter a Snowflake ID",
				Required:    true,
			},
		},
	}: Snowflake2time,
}

func Start() {
}

func Snowflake2time(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := i.ApplicationCommandData().Options[0].StringValue()

	Log.Verbose.Printf("Received Snowflake ID: %v", id)

	n := new(big.Int)
	n, ok := n.SetString(id, 10)
	if !ok {
		Log.Warn.Printf("Failed to convert string to bigint: %v", id)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Error: Failed to convert string to bigint. Please check your input.",
			},
		})
		return
	}

	timestamp := time.UnixMilli((n.Int64() >> 22) + 1420070400000)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("Timestamp: <t:%d> (Unix: `%d`)", timestamp.Unix(), timestamp.Unix()),
		},
	})
}
