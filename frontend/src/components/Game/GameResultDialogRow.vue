<script setup lang="ts">
import { ResultDealer, ResultPlayer } from "@/types/Player"
import Card from "@/components/Card.vue"

const props = defineProps<{
  isDealer: boolean
  player: ResultDealer | ResultPlayer
}>()
</script>

<template>
  <v-row align="center" dense justify="center">
    <v-col v-if="props.isDealer" cols="4" class="name">Dealer</v-col>
    <v-col v-else cols="2" class="name">{{
      (props.player as ResultPlayer).name
    }}</v-col>

    <v-col v-if="!props.isDealer" cols="2" class="chip">
      <v-chip
        v-if="(props.player as ResultPlayer).isWin"
        size="x-large"
        color="success"
        variant="elevated"
        >WIN</v-chip
      >
      <v-chip v-else size="x-large" color="error" variant="elevated"
        >LOSE</v-chip
      >
    </v-col>
    <v-col cols="1" class="score">{{ props.player.score }}</v-col>

    <v-col cols="5">
      <v-container>
        <v-row justify="center" align="center">
          <v-col v-for="card in props.player.hand" class="card">
            <Card :card="card" :size="4" />
          </v-col>
        </v-row>
      </v-container>
    </v-col>
  </v-row>
</template>

<style scoped lang="scss">
.row {
  margin-left: auto;
}
.name {
  font-size: 2rem;
  font-weight: bold;
}
.chip {
  text-align: center;
}

.score {
  font-size: 1.5rem;
  font-weight: bold;
  text-align: center;
}

.card {
  text-align: center;
}
</style>
