<script setup lang="ts">
import Card from "@/components/Card.vue"
import { Player, SuitIcon } from "@/types/Player"
import { defineProps } from "vue"
import { CardColor } from "@/utils/Color"

type Props = {
  player: Player & { icon: SuitIcon }
  hasTurn: boolean
}

const { player, hasTurn } = defineProps<Props>()
</script>

<template>
  <v-card :class="hasTurn && 'v-card--turn-holder'" :color="CardColor(player)">
    <template v-slot:title>
      <v-chip v-if="player.isBust" color="error" variant="elevated"
        >BUSTED</v-chip
      >
      <v-chip v-if="player.isStand" color="primary" variant="elevated"
        >STAND</v-chip
      >

      <v-icon size="x-large" color="secondary">{{ player.icon }}</v-icon>
      <span class="player-name">{{ player.name }}</span>
      <span class="player-score">score: {{ player.score }}</span>
    </template>

    <div class="hand">
      <span v-for="card in player.hand">
        <Card :card="card" :size="5" />
      </span>
    </div>
  </v-card>
</template>

<style scoped lang="scss">
.v-chip {
  margin-right: 1.5rem;
}
.v-card {
  padding-left: 1rem;
  padding-right: 1rem;
  padding-bottom: 0.3rem;

  &--turn-holder {
    border: solid 0.3rem rgb(var(--v-theme-primary));
  }
}
.hand {
  display: flex;
  align-items: center;
  *:nth-child(n) {
    margin-left: 0.2rem;
  }

  *:last-child {
    margin-right: 0rem;
  }
}

.player-name {
  font-size: 1.2rem;
  font-weight: bold;
}
.player-score {
  padding-left: 1.5rem;
  font-size: 1.2em;
  font-weight: bold;
}
</style>
