package main

func reverse(sentence string) string {
	list := []rune(sentence)
	for x, y := 0, len(sentence)-1; x < y; x, y = x+1, y-1 {
		list[x], list[y] = list[y], list[x]
	}

	return string(list)
}
