package converter

import (
	"casino/model/dealer"
	"casino/model/player"
	"casino/model/round"
)

type ResultRoundConverter struct {
	resultPlayerConverter Converter[player.Player, ResultPlayer]
	resultDealerConverter Converter[dealer.Dealer, ConvertedResultDealer]
}

func (rrc *ResultRoundConverter) Convert(currentRound *round.Round) *ConvertedResultRound {
	players := []*ResultPlayer{}
	for _, player := range currentRound.Players {
		players = append(players, rrc.resultPlayerConverter.Convert(player))
	}
	dealer := rrc.resultDealerConverter.Convert(currentRound.Dealer)

	return NewConvertedResultRound(players, dealer)
}

type ConvertedResultRound struct {
	Players []*ResultPlayer
	Dealer  *ConvertedResultDealer
}

func NewResultRoundConverter(
	rpc Converter[player.Player, ResultPlayer],
	rdc Converter[dealer.Dealer, ConvertedResultDealer],
) *ResultRoundConverter {
	return &ResultRoundConverter{
		resultPlayerConverter: rpc,
		resultDealerConverter: rdc,
	}
}

func NewConvertedResultRound(
	players []*ResultPlayer,
	dealer *ConvertedResultDealer,
) *ConvertedResultRound {
	return &ConvertedResultRound{
		Players: players,
		Dealer:  dealer,
	}
}
