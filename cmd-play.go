package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

func play(session *discordgo.Session, mCreate *discordgo.MessageCreate, guildID, channelID, url string) string {
	_, isPlaying := session.VoiceConnections[mCreate.GuildID]
	if isPlaying {
		return "Already playing."
	}
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

		embed := NewEmbed().
			SetTitle("Select a song.").
			AddField("1.", songNames[0]).
			AddField("2.", songNames[1]).
			AddField("3.", songNames[2]).
			SetColor(0xff0000).
			MessageEmbed

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

	encoded, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		return "Failed to encode audio."
	}
	defer encoded.Cleanup()

	embed := NewEmbed().
		SetTitle("Now playing "+audio.Title).
		AddField("Uploaded By", audio.Author).
		SetThumbnail(audio.GetThumbnailURL(ytdl.ThumbnailQualityHigh).String()).
		SetColor(0xff0000).
		MessageEmbed

	session.ChannelMessageSendEmbed(mCreate.ChannelID, embed)

	done := make(chan error)
	dca.NewStream(encoded, voice, done)
	<-done

	return "Done playing " + audio.Title

}
