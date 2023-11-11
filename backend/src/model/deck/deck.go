package deck

import (
	"casino/model/card"
	"casino/model/card/rank"
	"casino/model/card/suit"
	"math/rand"
)

type Deck struct {
	Cards []*card.Card
}

func NewDeck() *Deck {
	deck := Deck{}
	suits := [4]suit.Suit{suit.Hearts, suit.Diamonds, suit.Spades, suit.Clubs}
	ranks := [13]rank.Rank{rank.Ace, rank.Two, rank.Three, rank.Four, rank.Five, rank.Six, rank.Seven, rank.Eight, rank.Nine, rank.Ten, rank.Jack, rank.Queen, rank.King}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := card.NewCard(suit, rank)
			deck.Cards = append(deck.Cards, card)
		}
	}
	return &deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw() *card.Card {
	drawnCard := d.Cards[0]
	d.Cards = d.Cards[1:]
	return drawnCard
}

func (d *Deck) InitialHand() []*card.Card {
	drawnCard := []*card.Card{d.Cards[0], d.Cards[1]}
	d.Cards = d.Cards[2:]
	return drawnCard
}
