package card

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
)

type Card struct {
	Suit suit.Suit
	Rank rank.Rank
}

func NewCard(suit suit.Suit, rank rank.Rank) *Card {
	return &Card{suit, rank}
}
