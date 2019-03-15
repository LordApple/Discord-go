package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func pickle(users []*discordgo.User) string {
	exists := make(map[string]string)
	dongs := []string{}
	names := []string{}
	msg := ""

	for _, user := range users {
		if _, isExist := exists[user.Username]; isExist {
			continue
		}
		uID, _ := strconv.ParseInt(user.ID, 10, 64)
		rand.Seed(uID)

		dongs = append(dongs, fmt.Sprintf("8%vD", fmt.Sprintf(strings.Repeat("=", rand.Intn(30)))))
		names = append(names, user.Username)
		exists[user.Username] = ""
	}

	sort.Sort(byLength(dongs))

	for x := 0; x < len(names); x++ {
		msg += fmt.Sprintf("**%v's size:**\n%v\n", names[x], dongs[x])
	}

	return msg
}
