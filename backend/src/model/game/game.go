package game

import (
	"casino/model/deck"
	"casino/model/player"
	"errors"
	"strconv"
)

type Game struct {
	Deck          *deck.Deck
	CurrentPlayer player.PlayerId
	Players       []*player.Player
}

func CreateNewGame() *Game {
	game := Game{
		Deck: deck.NewDeck(),
	}
	return &game
}

func (g *Game) GetPlayer(player_id player.PlayerId) (*player.Player, error) {
	for _, player := range g.Players {
		if player.Id == player_id {
			return player, nil
		}
	}
	return nil, errors.New("PlayerId Not Found")
}

func (g *Game) GenerateUniquePlayerId() player.PlayerId {
	for i := 0; ; i++ {
		isUnique := true
		for _, p := range g.Players {
			if p.Id == player.PlayerId(strconv.Itoa(i)) {
				isUnique = false
				break
			}
		}
		if isUnique {
			return player.PlayerId(strconv.Itoa(i))
		}
	}
}

func (g *Game) Hit(id player.PlayerId) error {
	for _, p := range g.Players {
		if p.Id == id {
			if p.IsFold {
				return errors.New("player already folded")
			}
			p.Hand = append(p.Hand, g.Deck.Draw())
		}
	}
	return nil
}

func (g *Game) Fold(id player.PlayerId) error {
	for _, p := range g.Players {
		if p.Id == id {
			if p.IsFold {
				return errors.New("player already folded")
			}
			p.IsFold = true
		}
	}
	return nil
}
