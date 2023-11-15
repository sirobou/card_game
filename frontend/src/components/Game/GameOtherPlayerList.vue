<script setup lang="ts">
import GameOtherPlayerListItem from "@/components/Game/GameOtherPlayerListItem.vue"
import { computed } from "vue"
import { Player } from "@/types/Player"
import { playersWithIcon } from "@/utils/Player"

type Props = {
  players: Player[]
  currentPlayer: Player["id"]
}

const props = defineProps<Props>()
const playerWithIcons = computed(() => playersWithIcon(props.players))
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
