<script setup lang="ts">
import GameOtherPlayerListItem from "@/components/Game/GameOtherPlayerListItem.vue"
import { defineProps } from "vue"
import { Player, SuitIcon } from "@/types/Player"

type Props = {
  players: Player[]
  currentPlayer: Player["id"]
}

const { players } = defineProps<Props>()
const playerWithIcons = players.map<Player & { icon: SuitIcon }>(
  (player, i) => {
    const icon: SuitIcon[] = [
      "mdi-cards-playing-spade",
      "mdi-cards-playing-heart",
      "mdi-cards-playing-club",
      "mdi-cards-playing-diamond",
    ]

    return {
      ...player,
      icon: icon[i % icon.length],
    }
  }
)
</script>

<template>
  <div class="player-container">
    <span v-for="player in playerWithIcons" :key="player.id">
      <GameOtherPlayerListItem
        :player="player"
        :has-turn="player.id === currentPlayer"
      />
    </span>
  </div>
</template>

<style scoped lang="scss">
.player-container {
  height: 100%;
  max-width: 40rem;
  min-width: 20rem;
  *:nth-child(n) {
    margin-bottom: 0.5rem;
  }
}
</style>
