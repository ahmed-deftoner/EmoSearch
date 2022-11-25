package bot

import (
	"fmt"

	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/spotifytest/pkg/config"
	"github.com/spotifytest/pkg/parser"
)

var BotID string
var goBot *discordgo.Session
var msgs = []string{"uwu", "oni chan", "fuck me, baby", "cum inside me", "ara ara"}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content != "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, msgs[rand.Intn(5)])
		arr := parser.HandleRequest()
		if arr == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "empty arr")
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
