<script setup lang="ts">
import { defineProps, ref, onMounted } from "vue"
import GameResultDialogRow from "@/components/Game/GameResultDialogRow.vue"
import { BASE_URL } from "@/consts/const"
import { ResultDealer, ResultPlayer } from "@/types/Player"
import { toResultDealer, toResultPlayer } from "@/utils/Player"
import router from "@/router"

const props = defineProps<{
  dialog: boolean
}>()

const dealer = ref<ResultDealer | null>(null)
const players = ref<ResultPlayer[]>([])

onMounted(async () => {
  const resp = await fetch(`${BASE_URL}/api/round/result`)
  const res = await resp.json()
  dealer.value = toResultDealer(res.Dealer)
  players.value = res.Players.map(toResultPlayer)
})

const pushLobby = () => {
  router.push("/lobby")
}
</script>

<template>
  <v-dialog v-model="props.dialog" :max-width="'90vw'">
    <v-card>
      <v-container>
        <GameResultDialogRow v-if="dealer" :is-dealer="true" :player="dealer" />
        <GameResultDialogRow
          v-for="player in players"
          :key="player.id"
          :is-dealer="false"
          :player="player"
        />
      </v-container>
      <v-card-actions>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="pushLobby"
        >
          ロビーに戻る
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped lang="scss">
.v-card-actions {
  justify-content: center;
  margin-bottom: 0.5rem;
}
</style>
