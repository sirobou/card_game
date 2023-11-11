package converter

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
	"casino/model/player"
)

type PlayerConverter struct {
}

func (pc *PlayerConverter) Convert(originalPlayer *player.Player) *ConvertedPlayer {
	convertedHand := []*ConvertedCard{}
	for _, card := range originalPlayer.Hand {
		convertedHand = append(
			convertedHand,
			NewConvertedCard(suit.ConvertSuitToString(card.Suit), rank.ConvertRankToString(card.Rank)),
		)
	}
	return NewConvertedPlayer(originalPlayer.Id, originalPlayer.Name.Value, originalPlayer.InRound, originalPlayer.IsStand, convertedHand, originalPlayer.IsBust(), originalPlayer.GetHandTotal())
}

type ConvertedCard struct {
	Suit string
	Rank string
}

type ConvertedPlayer struct {
	Id        player.PlayerId
	Name      string
	InTable   bool
	IsStand   bool
	Hand      []*ConvertedCard
	IsBust    bool
	TotalHand int
}

func NewPlayerConverter() *PlayerConverter {
	return &PlayerConverter{}
}

func NewConvertedPlayer(
	id player.PlayerId,
	name string,
	inTable bool,
	isStand bool,
	hand []*ConvertedCard,
	isBust bool,
	totalHand int,
) *ConvertedPlayer {
	return &ConvertedPlayer{
		Name:      name,
		Id:        id,
		InTable:   inTable,
		IsStand:   isStand,
		Hand:      hand,
		IsBust:    isBust,
		TotalHand: totalHand,
	}
}

func NewConvertedCard(suit string, rank string) *ConvertedCard {
	return &ConvertedCard{
		suit,
		rank,
	}
}
