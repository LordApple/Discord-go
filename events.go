package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	split := strings.Split(m.Content, " ")
	cmd := split[0]

	//Bot won't respond to itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	//if message author is a bot, return
	if m.Author.Bot {
		return
	}

	if cmd == prefix+"about" {
		s.ChannelMessageSendEmbed(m.ChannelID, aboutBot(s))
	}
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateStreamingStatus(0, "nhentai.net", "https://www.twitch.tv/twitchbot_discord")
	fmt.Println("Logged in as " + r.User.Username + "#" + r.User.Discriminator + "  ID: " + r.User.ID)
	fmt.Println("--------")
	fmt.Printf("Current DiscordGO Version: %v | Current Golang Version: %v\n", discordgo.VERSION, runtime.Version())
	fmt.Println("--------")
	fmt.Printf("Use this link to invite %v:\n", r.User.Username+"#"+r.User.Discriminator)
	fmt.Printf("https://discordapp.com/oauth2/authorize?client_id=%v&scope=bot&permissions=8\n", r.User.ID)
	fmt.Println("--------")
	fmt.Println("Created by Apple#1337")

}
