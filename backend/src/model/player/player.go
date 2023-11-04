package player

import (
	"casino/model/card"
	"casino/model/card/rank"
	"casino/model/deck"
)

type Player struct {
	Id     PlayerId
	Name   PlayerName
	IsJoin bool
	IsFold bool
	Hand   []*card.Card
}

func NewPlayer(deck *deck.Deck, name string, id PlayerId) (*Player, error) {
	playerName, err := NewPlayerName(name)
	if err != nil {
		return nil, err
	}
	player := Player{
		Id:     id,
		Name:   playerName,
		IsJoin: false,
		IsFold: false,
		Hand:   deck.InitialHand(),
	}
	return &player, nil
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
