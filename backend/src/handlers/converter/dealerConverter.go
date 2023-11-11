package converter

import (
	"casino/model/card/rank"
	"casino/model/card/suit"
	"casino/model/dealer"
)

type DealerConverter struct {
}

func (dc *DealerConverter) Convert(originalDealer *dealer.Dealer) *ConvertedDealer {
	convertedCard := NewConvertedCard(suit.ConvertSuitToString(originalDealer.PublicHand().Suit), rank.ConvertRankToString(originalDealer.PublicHand().Rank))
	return NewConvertedDealer(convertedCard, len(originalDealer.Hand), originalDealer.IsBust(), originalDealer.IsStand())
}

type ConvertedDealer struct {
	PublicHand *ConvertedCard
	HandCount  int
	IsBust     bool
	IsStand    bool
}

func NewConvertedDealer(
	publicHand *ConvertedCard,
	handCount int,
	isBust bool,
	isStand bool,
) *ConvertedDealer {
	return &ConvertedDealer{
		PublicHand: publicHand,
		HandCount:  handCount,
		IsBust:     isBust,
		IsStand:    isStand,
	}
}

func NewDealerConverter() *DealerConverter {
	return &DealerConverter{}
}
