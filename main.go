package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string
)

func main() {
	prefix = os.Getenv("PREFIX")

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Failed to create session", err)
	}

	//Add all event handelers
	dg.AddHandler(onMessage)
	dg.AddHandler(onReady)

	if err := dg.Open(); err != nil {
		fmt.Println("Failed to create websocket", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

	dg.Close()
}
