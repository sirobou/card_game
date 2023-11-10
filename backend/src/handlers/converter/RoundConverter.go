package converter

import (
	"casino/model/deck"
	"casino/model/player"
	"casino/model/round"
)

type RoundConverter struct {
	playerConverter Converter[player.Player, ConvertedPlayer]
}

func (r *RoundConverter) Convert(currentRound *round.Round) *ConvertedRound {
	convertedPlayers := []*ConvertedPlayer{}
	for _, pl := range currentRound.Players {
		convertedPlayers = append(convertedPlayers, r.playerConverter.Convert(pl))
	}

	return NewConvertedRound(currentRound.Deck, currentRound.CurrentPlayer, convertedPlayers)
}

type ConvertedRound struct {
	IsEnd         bool
	CurrentPlayer player.PlayerId
	Players       []*ConvertedPlayer
}

func NewConvertedRound(deck *deck.Deck, currentPlayer player.PlayerId, players []*ConvertedPlayer) *ConvertedRound {
	return &ConvertedRound{
		IsEnd:         false,
		CurrentPlayer: currentPlayer,
		Players:       players,
	}
}

func NewRoundConverter(playerConverter Converter[player.Player, ConvertedPlayer]) *RoundConverter {
	return &RoundConverter{playerConverter}
}
