package main

import (
	"github.com/bwmarrin/discordgo"
)

func aboutBot(s *discordgo.Session) *discordgo.MessageEmbed {
	embed := NewEmbed().
		SetTitle("About "+s.State.User.Username).
		AddField("Name:", s.State.User.Username+"#"+s.State.User.Discriminator).
		SetThumbnail(s.State.User.AvatarURL("")).
		AddField("ID:", s.State.User.ID).
		MessageEmbed
	return embed
}
