package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

func play(session *discordgo.Session, mCreate *discordgo.MessageCreate, guildID, channelID, url string) string {
	_, isPlaying := session.VoiceConnections[mCreate.GuildID]
	if isPlaying {
		return "Already playing."
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

	format := audio.Formats[0]
	downloadURL, err := audio.GetDownloadURL(format)
	if err != nil {
		return "Failed to get download url."
	}

	encoded, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		return "Failed to encode audio."
	}
	defer encoded.Cleanup()

	session.ChannelMessageSend(mCreate.ChannelID, "Now playing "+audio.Title)

	done := make(chan error)
	dca.NewStream(encoded, voice, done)
	<-done

	return "Done playing " + audio.Title

}
