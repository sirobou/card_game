package converter

import (
	"casino/model/player"
	"casino/model/round"
)

type ResultPlayerConverter struct {
	playerConverter Converter[player.Player, ConvertedPlayer]
	round           *round.Round
}

func (rpc *ResultPlayerConverter) Convert(originalPlayer *player.Player) *ResultPlayer {
	convertedPlayer := rpc.playerConverter.Convert(originalPlayer)
	return NewResultPlayer(*convertedPlayer, rpc.round.IsWin(originalPlayer))
}

func NewResultPlayerConverter(
	playerConverter Converter[player.Player, ConvertedPlayer],
	round *round.Round,
) *ResultPlayerConverter {
	return &ResultPlayerConverter{
		playerConverter,
		round,
	}
}

type ResultPlayer struct {
	ConvertedPlayer
	IsWin bool
}

func NewResultPlayer(
	convertedPlayer ConvertedPlayer,
	isWin bool,
) *ResultPlayer {
	return &ResultPlayer{
		ConvertedPlayer: convertedPlayer,
		IsWin:           isWin,
	}
}
