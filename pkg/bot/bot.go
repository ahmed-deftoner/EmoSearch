package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spotifytest/pkg/config"
	"github.com/spotifytest/pkg/parser"
)

var BotID string
var goBot *discordgo.Session

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content != "" {
		if m.Content == "help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "input formar: <album/playlist name>-<type>-<emotion>-<number of songs>\n e.g Sempiternal-a-sad-5 (5 saddest songs from album with the name sempiternal)\nOptions\n type = a (album) | p (playlist)\nemotion = sad | happy | instrumental | vocal | acoustic | intense | dance")
			return
		}
		arr := parser.HandleRequest(m.Content)
		if arr == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "album/playlist not found")
			return
		}
		for i := 0; i < len(arr); i++ {
			_, _ = s.ChannelMessageSend(m.ChannelID, arr[i])
		}
	}
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("horny bot is running...")
}
