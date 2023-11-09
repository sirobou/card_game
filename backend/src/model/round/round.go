package round

import (
	"casino/model/deck"
	"casino/model/player"
	"errors"
)

type Round struct {
	Deck          *deck.Deck
	currentPlayer player.PlayerId
	Players       []*player.Player
}

func NewRound(players []*player.Player) *Round {
	round := Round{
		Deck: deck.NewDeck(),
		//TODO currentPlayer
		Players: players,
	}
	return &round
}

func (r *Round) GetPlayer(playerId player.PlayerId) (*player.Player, error) {
	for _, player := range r.Players {
		if player.Id == playerId {
			return player, nil
		}
	}
	return nil, errors.New("PlayerId Not Found")
}

func (r *Round) Hit(id player.PlayerId) error {
	for _, p := range r.Players {
		if p.Id == id {
			if p.IsFold {
				return errors.New("player already folded")
			}
			if p.IsBurst() {
				return errors.New("player already bursted")
			}
			p.Hand = append(p.Hand, r.Deck.Draw())
		}
	}
	return nil
}

func (r *Round) Fold(id player.PlayerId) error {
	for _, p := range r.Players {
		if p.Id == id {
			if p.IsFold {
				return errors.New("player already folded")
			}
			if p.IsBurst() {
				return errors.New("player already bursted")
			}
			p.IsFold = true
		}
	}
	return nil
}
