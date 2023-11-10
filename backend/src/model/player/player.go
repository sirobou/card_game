package player

import (
	"casino/consts"
	"casino/model/card"
	"casino/model/card/rank"
)

type Player struct {
	Id      PlayerId
	Name    PlayerName
	InRound bool
	IsFold  bool
	Hand    []*card.Card
}

func NewPlayer(name string, id PlayerId) (*Player, error) {
	playerName, err := NewPlayerName(name)
	if err != nil {
		return nil, err
	}
	player := Player{
		Id:      id,
		Name:    playerName,
		InRound: false,
		IsFold:  false,
		Hand:    []*card.Card{},
	}
	return &player, nil
}

// func NewPlayer(deck *deck.Deck, name string, id PlayerId) (*Player, error) {
// 	playerName, err := NewPlayerName(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	player := Player{
// 		Id:      id,
// 		Name:    playerName,
// 		InRound: false,
// 		IsFold:  false,
// 		Hand:    deck.InitialHand(),
// 	}
// 	return &player, nil
// }

func (p *Player) GetHandTotal() int {
	hand_total := 0
	for _, card := range p.Hand {
		switch card.Rank {
		case rank.Ten, rank.Jack, rank.Queen, rank.King:
			hand_total += 10
		case rank.Ace:
			if hand_total+11 > consts.BURST_THRESHHOLD {
				hand_total += 1
			} else {
				hand_total += 11
			}
		default:
			hand_total += int(card.Rank)
		}
	}
	return hand_total
}

func (p *Player) IsBurst() bool {
	return p.GetHandTotal() > consts.BURST_THRESHHOLD
}
