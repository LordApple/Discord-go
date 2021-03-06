package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	split := strings.Split(strings.ToLower(m.Content), " ")
	splitNormal := strings.Split(m.Content, " ")
	cmd := split[0]

	//if message author is a bot, return
	if m.Author.Bot {
		return
	}

	if cmd == prefix+"about" {
		s.ChannelMessageSendEmbed(m.ChannelID, aboutBot(s))
	}

	if cmd == prefix+"pickle" {
		users := m.Mentions
		if len(users) < 1 {
			users = append(users, m.Author)
		}
		s.ChannelMessageSend(m.ChannelID, pickle(users))
	}

	if cmd == prefix+"8ball" || cmd == prefix+"8" {
		question := strings.Join(split[1:], " ")
		s.ChannelMessageSendEmbed(m.ChannelID, _8ball(question))
	}

	if cmd == prefix+"echo" {
		echo := strings.Join(splitNormal[1:], " ")
		s.ChannelMessageSend(m.ChannelID, echo)
	}

	if cmd == prefix+"reverse" || cmd == prefix+"rev" {
		sentence := strings.Join(splitNormal[1:], " ")
		s.ChannelMessageSend(m.ChannelID, reverse(sentence))
	}

	if cmd == prefix+"greentext" {
		s.ChannelMessageSendEmbed(m.ChannelID, greentext())
	}

	if cmd == prefix+"play" {
		url := strings.Join(splitNormal[1:], " ")
		isVoice := false
		guild, _ := s.State.Guild(m.GuildID)

		if len(url) == 0 {
			s.ChannelMessageSend(m.ChannelID, "No URL found.")
			return
		}

		for _, user := range guild.VoiceStates {
			if user.UserID == m.Author.ID {
				isVoice = true
				s.ChannelMessageSend(m.ChannelID, play(s, m, m.GuildID, user.ChannelID, url))
				break
			} else {
				isVoice = false
			}
		}

		if !isVoice {
			s.ChannelMessageSend(m.ChannelID, "You must be in a voice channel.")
		}
	}
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	go changeStatus(s)
	fmt.Println("Logged in as " + r.User.Username + "#" + r.User.Discriminator + "  ID: " + r.User.ID)
	fmt.Println("--------")
	fmt.Printf("Current DiscordGO Version: %v | Current Golang Version: %v\n", discordgo.VERSION, runtime.Version())
	fmt.Println("--------")
	fmt.Printf("Use this link to invite %v:\n", r.User.Username+"#"+r.User.Discriminator)
	fmt.Printf("https://discordapp.com/oauth2/authorize?client_id=%v&scope=bot&permissions=8\n", r.User.ID)
	fmt.Println("--------")
	fmt.Println("Created by Apple#1337")

}

func changeStatus(s *discordgo.Session) {
	var cfg config
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to find config.json file")
		return
	}
	json.Unmarshal(bytes, &cfg)

	for {
		rand.Seed(time.Now().UnixNano())
		if cfg.Activity.Type == 3 {
			s.UpdateStreamingStatus(0, cfg.Activity.Gamelist[rand.Intn(len(cfg.Activity.Gamelist))], "https://www.twitch.tv/twitchbot_discord")
		}
		if cfg.Activity.Type == 2 {
			s.UpdateListeningStatus(cfg.Activity.Gamelist[rand.Intn(len(cfg.Activity.Gamelist))])
		}
		if cfg.Activity.Type == 1 {
			s.UpdateStatus(0, cfg.Activity.Gamelist[rand.Intn(len(cfg.Activity.Gamelist))])
		}
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill)
		time.Sleep(60 * time.Second)
	}
}
