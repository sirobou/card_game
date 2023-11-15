import {
  rankTable,
  RankKey,
  Rank,
  CardImage,
  Suit,
  Card,
  CardFromApi,
} from "../types/Card"

export const convertRank = (rank: RankKey): Rank => {
  return rankTable[rank]
}

export const convertCardImage = (card: Card): CardImage => {
  return `${card.suit.toLowerCase() as Lowercase<Suit>}_${
    card.rank.toLowerCase() as Lowercase<Rank>
  }`
}

export const toCard = (card: CardFromApi): Card => {
  return {
    rank: convertRank(card.Rank),
    suit: card.Suit.slice(0, -1) as Suit,
  }
}
