<script setup lang="ts">
import { Player, Dealer } from "@/types/Player"
import GamePlayer from "@/components/Game/GamePlayer.vue"
import GameOtherPlayerList from "@/components/Game/GameOtherPlayerList.vue"
import GameDealer from "@/components/Game/GameDealer.vue"

const player: Player = {
  id: 0,
  name: "player1",
  hand: [
    {
      suit: "Spade",
      rank: "Two",
    },
    {
      suit: "Spade",
      rank: "King",
    },
  ],
  score: 21,
  isStand: false,
  isBust: false,
}

const currentPlayer = 0

const dealer: Dealer = {
  publicHand: {
    suit: "Heart",
    rank: "Three",
  },
  handCount: 5,
  isStand: false,
}

const players = new Array(3).fill(null).map<Player>((_, i) => ({
  id: i + 1,
  name: `player${i}`,
  hand: [
    {
      suit: "Spade",
      rank: "Ace",
    },
    {
      suit: "Spade",
      rank: "King",
    },
    {
      suit: "Spade",
      rank: "Ace",
    },
    {
      suit: "Spade",
      rank: "King",
    },
  ],
  score: 21,
  isStand: false,
  isBust: false,
}))
</script>

<template>
  <v-container class="layout-root">
    <v-row>
      <v-col md="">
        <GameOtherPlayerList
          :players="players"
          :current-player="currentPlayer"
        />
      </v-col>

      <v-col>
        <GameDealer :dealer="dealer" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <GamePlayer :player="player" :has-turn="player.id === currentPlayer" />
    </v-row>
  </v-container>
</template>

<style scoped lang="scss">
.layout-root {
  margin-top: -4.5rem;
}
.game-conatiner {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;
}

.game-player {
  align-items: center;
}
</style>
