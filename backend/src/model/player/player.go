package player

import (
	"casino/model/card"
	"casino/model/card/rank"
	"casino/model/deck"
	"casino/model/player/id"
)

type Player struct {
	Id   id.PlayerId
	Hand []*card.Card
}

func NewPlayer(d *deck.Deck) Player {
	player := Player{Id: 1, Hand: d.InitialHand()}
	return player
}

func (p *Player) GetHandTotal() int {
	hand_total := 0
	for _, card := range p.Hand {
		switch card.Rank {
		case rank.Ten, rank.Jack, rank.Queen, rank.King:
			hand_total += 10
		case rank.Ace:
			if hand_total+10 < 22 {
				hand_total += 10
			} else {
				hand_total += 10
			}
		default:
			hand_total += int(card.Rank)
		}
	}
	return hand_total
}
