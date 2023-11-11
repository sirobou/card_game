<script setup lang="ts">
import Title from "@/components/Title.vue"
import LobbyButton from "@/components/Lobby/LobbyButton.vue"
import LobbyPlayerCard from "@/components/Lobby/LobbyPlayerCard.vue"

const suitIcons = [
  "mdi-cards-playing-spade",
  "mdi-cards-playing-heart",
  "mdi-cards-playing-diamond",
  "mdi-cards-playing-club",
]

const players = [
  {
    icon: suitIcons[0],
    name: "プレイヤー1",
  },
  {
    icon: suitIcons[1],
    name: "プレイヤー2",
  },
  {
    icon: suitIcons[2],
    name: "プレイヤー3",
  },
  {
    icon: suitIcons[3],
    name: "プレイヤー4",
  },
]
</script>

<template>
  <Title />
  <v-container>
    <v-row v-for="player in players" align="center" justify="center">
      <v-col class="player-card">
        <LobbyPlayerCard :name="player.name" :icon="player.icon" />
      </v-col>
    </v-row>
    <v-row>
      <v-col class="start-btn">
        <LobbyButton />
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped lang="scss">
.player-card {
  max-width: 20rem;
  min-width: 20rem;
}
.start-btn {
  text-align: center;
}
</style>
