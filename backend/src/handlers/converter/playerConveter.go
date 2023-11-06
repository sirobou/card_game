package converter

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
	"casino/model/player"
)

type PlayerConverter struct {
}

func (pc *PlayerConverter) Convert(originalPlayer *player.Player) *ConvertedPlayer {
	convertedHand := []ConvertedCard{}
	for _, card := range originalPlayer.Hand {
		convertedHand = append(
			convertedHand,
			ConvertedCard{
				suit.ConvertSuitToString(card.Suit),
				rank.ConvertRankToString(card.Rank),
			},
		)
	}
	return NewConvertedPlayer(originalPlayer.Id, originalPlayer.Name, originalPlayer.InTable, originalPlayer.IsFold, convertedHand)
}

type ConvertedCard struct {
	Suit string
	Rank string
}

type ConvertedPlayer struct {
	Id      player.PlayerId
	Name    player.PlayerName
	InTable bool
	IsFold  bool
	Hand    []ConvertedCard
}

func NewPlayerConverter() *PlayerConverter {
	return &PlayerConverter{}
}

func NewConvertedPlayer(
	id player.PlayerId,
	name player.PlayerName,
	inTable bool,
	isFold bool,
	hand []ConvertedCard,
) *ConvertedPlayer {
	return &ConvertedPlayer{
		Name:    name,
		Id:      id,
		InTable: inTable,
		IsFold:  isFold,
		Hand:    hand,
	}
}
