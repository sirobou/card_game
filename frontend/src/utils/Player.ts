import { Player, SuitIcon } from "@/types/Player"
import { Suit } from "@/types/Card"

export const convertToIcon = (suit: Suit): SuitIcon => {
  return `mdi-cards-playing-${suit.toLowerCase() as Lowercase<Suit>}`
}

export const findOwnPlayer = (players: Player[], id: number): Player => {
  const own = players.find((player) => player.id === id)
  if (!own) throw new Error("own player not found")
  return own
}
