package converter

import (
	"casino/model/dealer"
	"casino/model/player"
	"casino/model/round"
)

type RoundConverter struct {
	playerConverter Converter[player.Player, ConvertedPlayer]
	dealerConverter Converter[dealer.Dealer, ConvertedDealer]
}

func (r *RoundConverter) Convert(currentRound *round.Round) *ConvertedRound {
	convertedPlayers := []*ConvertedPlayer{}
	for _, pl := range currentRound.Players {
		convertedPlayers = append(convertedPlayers, r.playerConverter.Convert(pl))
	}
	convertedDealer := r.dealerConverter.Convert(currentRound.Dealer)

	return NewConvertedRound(currentRound.IsEndRound(), currentRound.CurrentPlayer, convertedPlayers, convertedDealer)
}

type ConvertedRound struct {
	IsEnd         bool
	CurrentPlayer player.PlayerId
	Players       []*ConvertedPlayer
	Dealer        *ConvertedDealer
}

func NewConvertedRound(isEnd bool, currentPlayer player.PlayerId, players []*ConvertedPlayer, dealer *ConvertedDealer) *ConvertedRound {
	return &ConvertedRound{
		IsEnd:         isEnd,
		CurrentPlayer: currentPlayer,
		Players:       players,
		Dealer:        dealer,
	}
}

func NewRoundConverter(playerConverter Converter[player.Player, ConvertedPlayer], dealerConverter Converter[dealer.Dealer, ConvertedDealer]) *RoundConverter {
	return &RoundConverter{playerConverter, dealerConverter}
}
