<script setup lang="ts">
import { ref, watch, onMounted, computed, onBeforeUnmount } from "vue"
import Title from "@/components/Title.vue"
import LobbyButton from "@/components/Lobby/LobbyButton.vue"
import LobbyPlayerCard from "@/components/Lobby/LobbyPlayerCard.vue"
import { BASE_URL } from "@/consts/const"
import { Player } from "@/types/Player"
import { playersWithIcon, toPlayer } from "@/utils/Player"
import router from "@/router"

const players = ref<Player[]>([])
const setIntervalID = ref<number | null>(null)

onMounted(async () => {
  // 0.5sごとにロビー状態をポーリング
  setIntervalID.value = window.setInterval(async () => {
    const resp = await fetch(`${BASE_URL}/api/lobby`)
    if (resp.ok) {
      const res = await resp.json()
      players.value = res.map(toPlayer)
    } else {
      console.error(resp)
    }
  }, 1000)
})

watch(players, (newPlayers, oldPlayers) => {
  if (oldPlayers.length > 0 && newPlayers.length < oldPlayers.length) {
    alert("ロビーがリセットされました")
    router.push("/")
  }
})

onBeforeUnmount(() => {
  if (setIntervalID.value) {
    window.clearInterval(setIntervalID.value)
  }
})

const viewPlayers = computed(() => playersWithIcon(players.value))

const startGame = async () => {
  const resp = await fetch(`${BASE_URL}/api/lobby/start`)
  if (resp.ok) {
    router.push("/game")
  } else {
    console.error(resp)
    alert("エラーが発生しました")
  }
}

const reset = async () => {
  await fetch(`${BASE_URL}/api/lobby/reset`)
  router.push("/")
}
</script>

<template>
  <Title />
  <v-container>
    <v-row align="center" justify="center">
      <v-btn rounded color="warning" @click="reset" size="small"
        >リセット</v-btn
      >
    </v-row>
    <v-row v-for="player in viewPlayers" align="center" justify="center">
      <v-col class="player-card">
        <LobbyPlayerCard :name="player.name" :icon="player.icon" />
      </v-col>
    </v-row>
    <v-row>
      <v-col class="start-btn" align-self="end">
        <LobbyButton @click="startGame" />
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
