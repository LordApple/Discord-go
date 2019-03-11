package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

var (
	queue     = make(map[string][]string)
	streaming = make(map[string]*dca.StreamingSession)
)

func play(session *discordgo.Session, mCreate *discordgo.MessageCreate, guildID, channelID, url string) string {
	if !strings.HasPrefix(url, "https://www.youtube.com") {
		results, err := findVideo(url)
		if err != nil {
			return "Failed to get results."
		}

		urls := results
		songNames := []string{}
		selected := false

		for _, url := range urls {
			audio, err := ytdl.GetVideoInfo(url)
			if err != nil {
				return "Failed to get video info."
			}

			songNames = append(songNames, audio.Title)
		}
		fields := []*discordgo.MessageEmbedField{}
		for x := 1; x <= len(urls); x++ {
			fields = append(fields, &discordgo.MessageEmbedField{Name: strconv.FormatInt(int64(x), 10), Value: songNames[x-1]})
		}

		embed := NewEmbed().
			SetTitle("Select a song.").
			SetColor(0xff0000).
			MessageEmbed
		embed.Fields = fields

		session.ChannelMessageSendEmbed(mCreate.ChannelID, embed)

		waitFor := session.AddHandler(func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
			if msg.Author.ID == mCreate.Author.ID {
				switch msg.Content {
				case "1":
					selected = true
					url = urls[0]
				case "2":
					selected = true
					url = urls[1]
				case "3":
					selected = true
					url = urls[2]
				}
			}
		})
		time.Sleep(5 * time.Second)
		waitFor()
		if !selected {
			return "No song selected"
		}
		if streaming[mCreate.GuildID] != nil {
			streaming, _ := streaming[mCreate.GuildID].Finished()
			if !streaming {
				queue[mCreate.GuildID] = append(queue[mCreate.GuildID], url)
				return "Added song to queue."
			}
		}
	}
	if streaming[mCreate.GuildID] != nil {
		streaming, _ := streaming[mCreate.GuildID].Finished()
		if !streaming {
			queue[mCreate.GuildID] = append(queue[mCreate.GuildID], url)
			return "Added song to queue."
		}
	}

	voice, _ := session.ChannelVoiceJoin(guildID, channelID, false, true)
	defer voice.Disconnect()

	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 320
	options.Application = "lowdelay"

	audio, err := ytdl.GetVideoInfo(url)
	if err != nil {
		return "Failed to get video info."
	}

	downloadURL, err := audio.GetDownloadURL(audio.Formats[0])
	if err != nil {
		return "Failed to get download url."
	}

	encoding, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		return "Failed to encode audio."
	}
	defer encoding.Cleanup()

	embed := NewEmbed().
		SetTitle("Now playing "+audio.Title).
		AddField("Uploaded By", audio.Author).
		SetThumbnail(audio.GetThumbnailURL(ytdl.ThumbnailQualityHigh).String()).
		SetColor(0xff0000).
		MessageEmbed

	session.ChannelMessageSendEmbed(mCreate.ChannelID, embed)

	done := make(chan error)
	streaming[mCreate.GuildID] = dca.NewStream(encoding, voice, done)
	<-done

	for len(queue[mCreate.GuildID]) > 0 {
		songURL := queue[mCreate.GuildID][0]
		queue[mCreate.GuildID] = queue[mCreate.GuildID][1:]
		play(session, mCreate, mCreate.GuildID, mCreate.ChannelID, songURL)
	}

	return "Done playing"

}
