package main

type config struct {
	Token    string `json:"TOKEN"`
	Prefix   string `json:"PREFIX"`
	Activity struct {
		Gamelist []string `json:"GAMES"`
		Type     int      `json:"TYPE"`
	} `json:"ACTIVITY"`
}
