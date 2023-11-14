import { Player, SuitIcon } from "@/types/Player"
import { Suit } from "@/types/Card"

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
