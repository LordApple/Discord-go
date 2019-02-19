package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string
)

func main() {
	var cfg config
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to find config.json file")
		return
	}
	json.Unmarshal(bytes, &cfg)

	prefix = cfg.Prefix

	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Failed to create session", err)
	}

	//Add all event handelers
	dg.AddHandler(onMessage)
	dg.AddHandler(onReady)

	if err := dg.Open(); err != nil {
		fmt.Println("Failed to create websocket", err)
	}

	ch := make(chan int)
	<-ch
}
