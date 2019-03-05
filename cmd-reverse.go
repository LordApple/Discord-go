package main

import (
	"regexp"
	"strings"
)

func reverse(sentence string) string {
	reg, _ := regexp.Compile("[^a-zA-Z^<>{}\"/|;:.,~!?@#$%^=&*\\\\\\()\\[¿§«»ω⊙¤°℃℉€¥£¢¡®©0-9_+*$]")
	processedString := reg.ReplaceAllString(sentence, "")

	list := []rune(processedString)
	for x, y := 0, len(processedString)-1; x < y; x, y = x+1, y-1 {
		list[x], list[y] = list[y], list[x]
	}

	rev := strings.Replace(strings.Replace(string(list), "@", "@\u200B", -1), "&", "&\u200B", -1)
	return rev
}
