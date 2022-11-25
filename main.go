package main

import (
	"fmt"

	"github.com/spotifytest/pkg/bot"
	"github.com/spotifytest/pkg/config"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
