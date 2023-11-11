package round

import (
	"casino/model/dealer"
	"casino/model/deck"
	"casino/model/player"
	"casino/model/round/action"
	"errors"
)

type Round struct {
	Deck          *deck.Deck
	CurrentPlayer player.PlayerId
	Players       []*player.Player
	Dealer        *dealer.Dealer
}

func NewRound(players []*player.Player) *Round {
	round := Round{
		Deck:          deck.NewDeck(),
		CurrentPlayer: players[0].Id,
		Players:       players,
		Dealer:        dealer.NewDealer(),
	}
	round.Deck.Shuffle()
	round.DealCard()
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

func (r *Round) Action(id player.PlayerId, receivedAction action.Action) error {
	if r.IsEndRound() {
		return errors.New("round is over")
	}

	for _, p := range r.Players {
		if p.Id == id {
			if p.IsStand {
				return errors.New("player already standed")
			}
			if p.IsBust() {
				return errors.New("player already busted")
			}
			switch receivedAction {
			case action.Hit:
				p.Hand = append(p.Hand, r.Deck.Draw())
				break
			case action.Stand:
				p.IsStand = true
				break
			default:
				panic("👍🏻receive invalied Action👍🏻")
			}
		}
	}
	r.RotateTurn()
	return nil
}

func (r *Round) IsEndRound() bool {
	return len(r.actionablePlayers()) == 0
}

func (r *Round) DealCard() {
	r.Dealer.Hand = r.Deck.InitialHand()
	for _, player := range r.Players {
		player.Hand = r.Deck.InitialHand()
	}
}

func (r *Round) RotateTurn() {
	players := r.Players[:]

	currentPlayerIndex := 0
	for i := 0; i < len(players); i++ {
		if players[i].Id == r.CurrentPlayer {
			currentPlayerIndex = i
			break
		}
	}
	nextPlayerIndex := currentPlayerIndex + 1
	if nextPlayerIndex > len(players)-1 {
		nextPlayerIndex = 0
		r.dealerAction()
	}

	// currentPlayerの次のターンを持つアクション可能なプレイヤーの探索
	for _, player := range players[nextPlayerIndex:] {
		if !(player.IsStand || player.IsBust()) {
			r.CurrentPlayer = player.Id
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

func (r *Round) actionablePlayers() []*player.Player {
	players := []*player.Player{}
	for _, player := range r.Players {
		if player.IsStand || player.IsBust() {
			continue
		}
		players = append(players, player)
	}
	return players
}

func (r *Round) dealerAction() {
	if r.Dealer.IsStand() || r.Dealer.IsBust() {
		return
	}
	r.Dealer.Hand = append(r.Dealer.Hand, r.Deck.Draw())
}

func (r *Round) IsWin(player *player.Player) bool {
	if player.IsBust() {
		return false
	}
	if r.Dealer.IsBust() {
		return true
	}
	return player.GetHandTotal() >= r.Dealer.GetHandTotal()
}
