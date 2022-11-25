package bot

import (
	"fmt"
	"math/rand"

	"github.com/ahmed-deftoner/discord-bot/pkg/config"
	"github.com/bwmarrin/discordgo"
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
