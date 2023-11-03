package deck

import (
	"math/rand"
)

var suits = []string{"Hearts", "Diamonds", "Spades", "Clubs"}
var ranks = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	deck := Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Suit: suit, Rank: rank}
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw() Card {
	drawnCard := d.Cards[0]
	d.Cards = d.Cards[1:]
	return drawnCard
}
