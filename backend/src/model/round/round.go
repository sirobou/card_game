package round

import (
	"casino/model/deck"
	"casino/model/player"
	"errors"
	"fmt"
)

type Round struct {
	Deck          *deck.Deck
	CurrentPlayer player.PlayerId
	Players       []*player.Player
}

func NewRound(players []*player.Player) *Round {
	fmt.Println(players)
	round := Round{
		Deck:          deck.NewDeck(),
		CurrentPlayer: players[0].Id,
		Players:       players,
	}
	round.Deck.Shuffle()
	round.DealCard()
	return &round
}

func (r *Round) GetPlayer(playerId player.PlayerId) (*player.Player, error) {
	fmt.Println("works")
	fmt.Println(r.Players)
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
			r.Rotate()
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
			r.Rotate()
		}
	}
	return nil
}

func (r *Round) DealCard() {
	for _, player := range r.Players {
		player.Hand = r.Deck.InitialHand()
	}
}

func (r *Round) Rotate() {
	for i, player := range r.Players {
		if i+1 == len(r.Players) {
			r.CurrentPlayer = r.Players[0].Id
			// r.isNotBurstOrFold(r.Players[0])
			return
		}
		if player.Id == r.CurrentPlayer {
			r.CurrentPlayer = r.Players[i+1].Id
			// r.isNotBurstOrFold(r.Players[i+1])
			return
		}
	}
}

func (r *Round) IsPlayerTurn(playerId player.PlayerId) error {
	if r.CurrentPlayer != playerId {
		return errors.New("Incorrect player is attempting action")
	}
	return nil
}

func (r *Round) isNotBurstOrFold(p *player.Player) {
	if p.IsFold == true || p.IsBurst() == true {
		r.Rotate()
	}
}
