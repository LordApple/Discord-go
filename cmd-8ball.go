package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func _8ball(question string) *discordgo.MessageEmbed {
	var embed *discordgo.MessageEmbed
	answers := []string{"As I see it, yes", "It is certain", "It is decidedly so", "Most likely", "Outlook good",
		"Signs point to yes", "Without a doubt", "Yes", "Yes – definitely", "You may rely on it", "Reply hazy, try again",
		"Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again",
		"Don't count on it", "My reply is no", "My sources say no", "Outlook not so good", "Very doubtful"}

	if strings.HasSuffix(question, "?") {
		rand.Seed(time.Now().UnixNano())
		answer := answers[rand.Intn(len(answers))]

		embed = NewEmbed().
			SetTitle("8Ball response").
			AddField("Question:", question).
			AddField("Answer", answer).
			SetColor(0x10d16d).
			MessageEmbed

	} else {
		embed = NewEmbed().
			SetTitle("Error").
			SetDescription("Question must end in a question mark").
			SetColor(0xfc0509).
			MessageEmbed
	}

	return embed
}
