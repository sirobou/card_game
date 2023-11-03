package card

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
)

type Card struct {
	Suit suit.Suit
	Rank rank.Rank
}

type ConvertedCard struct {
	Suit string
	Rank string
}

func (c *Card) Convert() ConvertedCard {
	return ConvertedCard{
		suit.ConvertToString(c.Suit),
		rank.ConvertRankToString(c.Rank),
	}
}

func NewCard(suit suit.Suit, rank rank.Rank) *Card {
	return &Card{suit, rank}
}
