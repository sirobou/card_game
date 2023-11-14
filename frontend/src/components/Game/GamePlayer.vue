<script setup lang="ts">
import { Player } from "@/types/Player"
import GamePlayerActionForm from "@/components/Game/GamePlayerActionForm.vue"
import GamePlayerHand from "@/components/Game/GamePlayerHand.vue"
import { CardColor } from "@/utils/Color"

type Props = {
  player: Player
  hasTurn: boolean
}

const props = defineProps<Props>()
type Emit = {
  action: [param: "hit" | "stand"]
}

const emit = defineEmits<Emit>()
const hit = () => emit("action", "hit")
const stand = () => emit("action", "stand")
</script>

<template>
  <v-card
    :class="props.hasTurn && 'v-card--turn-holder'"
    :color="CardColor(props.player)"
  >
    <v-container>
      <v-row align="center" justify="center">
        <v-chip v-if="props.player.isBust" color="error" variant="elevated"
          >BUSTED</v-chip
        >
        <v-chip v-if="props.player.isStand" color="primary" variant="elevated"
          >STAND</v-chip
        >
        <span class="player-name">{{ props.player.name }}</span>
        <span class="player-score">score: {{ props.player.score }}</span>
      </v-row>
      <v-row align="center" justify="center">
        <GamePlayerActionForm
          @hit="hit"
          @stand="stand"
          :disabled="props.player.isStand || props.player.isBust"
        />
      </v-row>
      <v-row align="center" justify="center">
        <GamePlayerHand :hand="props.player.hand" />
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped lang="scss">
.v-chip {
  margin-right: 1rem;
}
.v-card {
  // max-width: 75rem;
  min-width: 100%;
  &--turn-holder {
    border: solid 0.3rem rgb(var(--v-theme-primary));
  }
}
.v-container {
  * {
    .v-row:nth-child(2) {
      margin-bottom: 0.1rem;
    }
  }
}

.player-name {
  font-size: 1.3rem;
  font-weight: bold;
}
.player-score {
  padding-left: 1.5rem;
  font-size: 1.3em;
  font-weight: bold;
}
</style>
