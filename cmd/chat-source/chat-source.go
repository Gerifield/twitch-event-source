package main

import (
	"flag"
	"log"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"

	"github.com/gerifield/twitch-event-source/model"
	"github.com/gerifield/twitch-event-source/queue"
	"github.com/gerifield/twitch-event-source/token"
)

func main() {
	channelName := flag.String("channel", "gerifield", "Twitch channel name (use , to separate multiple channel ids)")
	botName := flag.String("botName", "CoderBot42", "Bot name")
	clientID := flag.String("clientID", "", "Twitch App ClientID")
	clientSecret := flag.String("clientSecret", "", "Twitch App clientSecret")

	flag.Parse()

	tl := token.New(*clientID, *clientSecret)
	log.Println("Fetching token")
	token, err := tl.Get()
	if err != nil {
		log.Println(err)
		return
	}

	q := queue.New()

	client := twitch.NewClient(*botName, "oauth:"+token.AccessToken)

	client.OnPrivateMessage(func(m twitch.PrivateMessage) {
		cm := model.EventFrame{
			ChannelID: m.RoomID,
			Type:      model.ChatMessage,
			Payload:   m,
		}

		_ = q.Add(cm)
	})

	client.OnUserNoticeMessage(func(m twitch.UserNoticeMessage) {
		cm := model.EventFrame{
			ChannelID: m.RoomID,
			Type:      model.ChatMessage,
			Payload:   m,
		}

		_ = q.Add(cm)
	})

	client.Join(strings.Split(*channelName, ",")...)

	log.Println("Connect with client")
	err = client.Connect()
	if err != nil {
		log.Println(err)
	}
}
