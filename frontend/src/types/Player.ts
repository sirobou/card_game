import { Card, Suit } from "@/types/Card"

export type SuitIcon = `mdi-cards-playing-${Lowercase<Suit>}`

export type Player = {
  id: number
  name: string
  hand: Card[]
  score: number
  isStand: boolean
  isBust: boolean
}

export type Dealer = {
  publicHand: Card
  handCount: number
  isStand: boolean
} & Pick<Partial<Player>, "hand" | "score" | "isBust">
