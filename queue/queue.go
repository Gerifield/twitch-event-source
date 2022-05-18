package queue

import (
	"log"

	"github.com/gerifield/twitch-event-source/model"
)

// Logic .
type Logic struct {
}

func New() *Logic {
	return &Logic{}
}

func (l *Logic) Add(msg model.EventFrame) error {
	log.Printf("Message added to the queue: %v\n", msg)
	return nil
}
