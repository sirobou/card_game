export const rankTable = {
  Ace: "Ace",
  "2": "Two",
  "3": "Three",
  "4": "Four",
  "5": "Five",
  "6": "Six",
  "7": "Seven",
  "8": "Eight",
  "9": "Nine",
  "10": "Ten",
  Jack: "Jack",
  Queen: "Queen",
  King: "King",
} as const

export type RankKey = keyof typeof rankTable
export type Rank = (typeof rankTable)[RankKey]
export type Suit = "Spade" | "Heart" | "Club" | "Diamond"

export type CardImage = `${Lowercase<Suit>}_${Lowercase<Rank>}`

export type Card = {
  rank: Rank
  suit: Suit
}

export type CardFromApi = {
  Rank: RankKey
  Suit: `${Suit}s`
}
