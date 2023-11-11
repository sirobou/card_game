package dealer

import (
	"casino/consts"
	"casino/model/card"
	"casino/model/card/rank"
)

type Dealer struct {
	Hand []*card.Card
}

func NewDealer() *Dealer {
	dealer := Dealer{
		Hand: []*card.Card{},
	}
	return &dealer
}

func (d *Dealer) GetHandTotal() int {
	hand_total := 0
	for _, card := range d.Hand {
		switch card.Rank {
		case rank.Ten, rank.Jack, rank.Queen, rank.King:
			hand_total += 10
		case rank.Ace:
			if hand_total+11 > consts.BUST_THRESHHOLD {
				hand_total += 1
			} else {
				hand_total += 11
			}
		default:
			hand_total += int(card.Rank)
		}
	}
	return hand_total
}

func (d *Dealer) IsBust() bool {
	return d.GetHandTotal() > consts.BUST_THRESHHOLD
}

func (d *Dealer) IsStand() bool {
	return d.GetHandTotal() >= consts.DEALER_STAND_THRESHHOLD
}

func (d *Dealer) PublicHand() *card.Card {
	return d.Hand[0]
}
