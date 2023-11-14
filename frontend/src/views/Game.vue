<script setup lang="ts">
import { ref, onMounted, computed, onBeforeUnmount } from "vue"
import { Player, Dealer } from "@/types/Player"
import { toPlayer, toDealer } from "@/utils/Player"
import GamePlayer from "@/components/Game/GamePlayer.vue"
import GameOtherPlayerList from "@/components/Game/GameOtherPlayerList.vue"
import GameDealer from "@/components/Game/GameDealer.vue"
import GameResultDialog from "@/components/Game/GameResultDialog.vue"
import { BASE_URL } from "@/consts/const"

const players = ref<Player[]>([])
const isEnd = ref<boolean>(false)

const otherPlayers = computed(() =>
  players.value.filter<Player>(
    (p): p is Player => p.id !== localStorage.getItem("id")
  )
)
const player = computed(() =>
  players.value.find((p) => p.id === localStorage.getItem("id"))
)

const setIntervalId = ref<number | null>(null)

const currentPlayerId = ref<string>("")
const dealer = ref<Dealer | null>(null)
onMounted(async () => {
  setIntervalId.value = window.setInterval(async () => {
    const resp = await fetch(`${BASE_URL}/api/round`)
    if (resp.ok) {
      const res = await resp.json()
      players.value = res.Players.map(toPlayer)
      currentPlayerId.value = res.CurrentPlayer
      dealer.value = toDealer(res.Dealer)
      isEnd.value = res.IsEnd
    } else {
      console.log(resp)
    }
  }, 1000)
})

const action = async (param: "hit" | "stand") => {
  const resp = await fetch(`${BASE_URL}/api/round/action`, {
    method: "POST",
    body: JSON.stringify({ action: param, id: localStorage.getItem("id") }),
  })

  if (resp.ok) {
    const res = await resp.json()
    players.value = res.Players.map(toPlayer)
    currentPlayerId.value = res.CurrentPlayer
    dealer.value = toDealer(res.Dealer)
  } else {
    console.log(resp)
  }
}

onBeforeUnmount(() => {
  if (setIntervalId.value) {
    window.clearInterval(setIntervalId.value)
  }
})
</script>

<template>
  <GameResultDialog :dialog="isEnd" />
  <v-container class="layout-root">
    <v-row>
      <v-col>
        <GameOtherPlayerList
          :players="otherPlayers"
          :current-player="currentPlayerId"
        />
      </v-col>

      <v-col>
        <GameDealer v-if="dealer" :dealer="dealer" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <GamePlayer
        v-if="player"
        :player="player"
        :has-turn="player.id === currentPlayerId"
        @action="action"
      />
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
