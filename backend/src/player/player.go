package player

import (
	"casino/deck"
)

type Player struct {
	id   string
	Hand []deck.Card
}

func NewPlayer() Player {
	player := Player{id: "guys"}
	return player
}
