import { Card, Suit, CardFromApi } from "@/types/Card"

export type SuitIcon = `mdi-cards-playing-${Lowercase<Suit>}`

export type PlayerFromApi = {
  Id: string
  Name: string
  Hand: CardFromApi[]
  Score: number
  IsStand: boolean
  IsBust: boolean
  TotalHand: number
}

export type Player = {
  id: string
  name: string
  hand: Card[]
  score: number
  isStand: boolean
  isBust: boolean
}

export type DealerFromApi = {
  PublicHand: CardFromApi
  HandCount: number
  IsStand: boolean
}

export type Dealer = {
  publicHand: Card
  handCount: number
  isStand: boolean
}

export type ResultPlayerFromApi = PlayerFromApi & { IsWin: boolean }
export type ResultPlayer = Player & { isWin: boolean }

export type ResultDealerFromApi = {
  Hand: CardFromApi[]
  IsBust: boolean
  TotalHand: number
}
export type ResultDealer = {
  hand: Card[]
  isBust: boolean
  score: number
}
