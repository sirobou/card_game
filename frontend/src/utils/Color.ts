import { Player, Dealer } from "@/types/Player"

export const CardColor = (player: Player | Dealer) => {
  if (player.hasOwnProperty("isBust") && (player as Player).isBust) {
    return "#FFCDD2"
  }

  if (player.isStand) {
    return "#BBDEFB"
  }

  return ""
}
