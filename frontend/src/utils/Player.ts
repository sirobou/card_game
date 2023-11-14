import {
  Player,
  SuitIcon,
  PlayerFromApi,
  DealerFromApi,
  Dealer,
  ResultPlayerFromApi,
  ResultPlayer,
  ResultDealerFromApi,
  ResultDealer,
} from "@/types/Player"
import { Suit } from "@/types/Card"
import { toCard } from "@/utils/Card"

export const convertToIcon = (suit: Suit): SuitIcon => {
  return `mdi-cards-playing-${suit.toLowerCase() as Lowercase<Suit>}`
}

export const findOwnPlayer = (players: Player[], id: string): Player => {
  const own = players.find((player) => player.id === id)
  if (!own) throw new Error("own player not found")
  return own
}

export const playersWithIcon = (
  players: Player[]
): (Player & { icon: SuitIcon })[] => {
  const suitIcons: SuitIcon[] = [
    "mdi-cards-playing-spade",
    "mdi-cards-playing-heart",
    "mdi-cards-playing-diamond",
    "mdi-cards-playing-club",
  ]

  return players.map((player, index) => {
    return {
      ...player,
      icon: suitIcons[index % 4],
    }
  })
}

export const toPlayer = (playerFromApi: PlayerFromApi): Player => {
  return {
    id: playerFromApi.Id,
    name: playerFromApi.Name,
    hand: playerFromApi.Hand.map((card) => toCard(card)),
    score: playerFromApi.TotalHand,
    isStand: playerFromApi.IsStand,
    isBust: playerFromApi.IsBust,
  }
}

export const toDealer = (DealerFromApi: DealerFromApi): Dealer => {
  return {
    publicHand: toCard(DealerFromApi.PublicHand),
    handCount: DealerFromApi.HandCount,
    isStand: DealerFromApi.IsStand,
  }
}

export const toResultPlayer = (
  resultPlayerFromApi: ResultPlayerFromApi
): ResultPlayer => {
  return {
    ...toPlayer(resultPlayerFromApi),
    isWin: resultPlayerFromApi.IsWin,
  }
}

export const toResultDealer = (
  resultDealerFromApi: ResultDealerFromApi
): ResultDealer => {
  return {
    hand: resultDealerFromApi.Hand.map((card) => toCard(card)),
    isBust: resultDealerFromApi.IsBust,
    score: resultDealerFromApi.TotalHand,
  }
}
