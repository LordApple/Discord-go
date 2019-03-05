package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func pickle(users []*discordgo.User) string {
	dongs := make(map[string]string)
	msg := ""

	for _, user := range users {
		if _, isExist := dongs[user.Username]; isExist {
			continue
		}
		uID, _ := strconv.ParseInt(user.ID, 10, 64)
		rand.Seed(uID)
		dongs[user.Username] = fmt.Sprintf("8%vD", strings.Repeat("=", rand.Intn(30)))
	}

	for name, dong := range dongs {
		msg += fmt.Sprintf("**%v's size:**\n%v\n", name, dong)
	}

	return msg
}
