package converter

import (
	"casino/model/deck"
	"casino/model/game"
	"casino/model/player"
)

type GameConverter struct {
	playerConverter Converter[player.Player, ConvertedPlayer]
}

func (g *GameConverter) Convert(game *game.Game) *ConvertedGame {
	convertedPlayers := []*ConvertedPlayer{}
	for _, pl := range game.Players {
		convertedPlayers = append(convertedPlayers, g.playerConverter.Convert(pl))
	}

	return NewConvertedGame(game.Deck, game.CurrentPlayer, convertedPlayers)
}

type ConvertedGame struct {
	IsEnd         bool
	CurrentPlayer player.PlayerId
	Players       []*ConvertedPlayer
}

func NewConvertedGame(deck *deck.Deck, currentPlayer player.PlayerId, players []*ConvertedPlayer) *ConvertedGame {
	return &ConvertedGame{
		IsEnd:         false,
		CurrentPlayer: currentPlayer,
		Players:       players,
	}
}

func NewGameConverter(playerConverter Converter[player.Player, ConvertedPlayer]) *GameConverter {
	return &GameConverter{playerConverter}
}
