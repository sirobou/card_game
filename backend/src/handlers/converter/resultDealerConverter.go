package converter

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
	"casino/model/dealer"
)

type ResultDealerConverter struct{}

func (rdc *ResultDealerConverter) Convert(originalDealer *dealer.Dealer) *ConvertedResultDealer {
	convertedHand := []*ConvertedCard{}
	for _, card := range originalDealer.Hand {
		convertedHand = append(
			convertedHand,
			NewConvertedCard(suit.ConvertSuitToString(card.Suit), rank.ConvertRankToString(card.Rank)),
		)
	}
	return NewConvertedResultDealer(convertedHand, originalDealer.IsBust(), originalDealer.IsStand(), originalDealer.GetHandTotal())
}

type ConvertedResultDealer struct {
	Hand      []*ConvertedCard
	IsBust    bool
	IsStand   bool
	TotalHand int
}

func NewConvertedResultDealer(
	hand []*ConvertedCard,
	isBust bool,
	isStand bool,
	totalHand int,

) *ConvertedResultDealer {
	return &ConvertedResultDealer{
		Hand:      hand,
		IsBust:    isBust,
		IsStand:   isStand,
		TotalHand: totalHand,
	}
}

func NewReulstDealerConverter() *ResultDealerConverter {
	return &ResultDealerConverter{}
}
