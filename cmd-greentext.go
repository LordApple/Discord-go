package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func greentext() *discordgo.MessageEmbed {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.reddit.com/r/greentext/random", nil)
	if err != nil {
		embed := NewEmbed().
			SetTitle("Error").
			SetDescription("Failed to connect to reddit api").
			SetColor(0xfc0509).
			MessageEmbed
		return embed
	}
	req.Header.Set("User-Agent", "Discord-Bot by Apple#1337")

	resp, _ := client.Do(req)
	bytes, _ := ioutil.ReadAll(resp.Body)

	var redditAPI []reddit
	json.Unmarshal(bytes, &redditAPI)

	embed := NewEmbed().
		SetTitle(">greentext").
		SetImage(redditAPI[0].Data.Children[0].Data.URL).
		SetColor(0x10d16d).
		MessageEmbed
	return embed
}
