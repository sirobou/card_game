import { rankTable, RankKey, Rank, CardImage, Suit, Card } from "../types/Card"

export const convertRank = (rank: RankKey): Rank => {
  return rankTable[rank]
}

export const convertCardImage = (card: Card): CardImage => {
  return `${card.suit.toLowerCase() as Lowercase<Suit>}_${card.rank.toLowerCase() as Lowercase<Rank>}.svg`
}
