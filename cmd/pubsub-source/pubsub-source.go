package main

import (
	"flag"
	"log"

	"github.com/kelr/gundyr/pubsub"

	"github.com/gerifield/twitch-event-source/model"
	"github.com/gerifield/twitch-event-source/queue"
	"github.com/gerifield/twitch-event-source/token"
)

func main() {
	channelID := flag.String("channelID", "gerifield", "Twitch channelID")
	clientID := flag.String("clientID", "", "Twitch App ClientID")
	clientSecret := flag.String("clientSecret", "", "Twitch App clientSecret")

	flag.Parse()

	tl := token.New(*clientID, *clientSecret)
	log.Println("Fetching token")
	token, err := tl.Get([]string{"bits:read", "channel:read:redemptions"})
	if err != nil {
		log.Println(err)
		return
	}

	q := queue.New()

	client := pubsub.NewClient(*channelID, token)
	client.ListenChannelPoints(func(m *pubsub.ChannelPointsData) {
		cm := model.EventFrame{
			ChannelID: m.Redemption.ChannelID,
			Type:      model.ChannelPointRedeem,
			Payload:   m,
		}

		_ = q.Add(cm)
	})

	err = client.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	select {}
}
