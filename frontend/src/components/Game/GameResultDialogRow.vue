<script setup lang="ts">
import { defineProps } from "vue"
import { ResultDealer, ResultPlayer } from "@/types/Player"
import Card from "@/components/Card.vue"

const props = defineProps<{
  isDealer: boolean
  player: ResultDealer | ResultPlayer
}>()
</script>

<template>
  <v-row align="center">
    <v-col v-if="!props.isDealer">
      <v-chip
        v-if="(props.player as ResultPlayer).isWin"
        color="success"
        variant="elevated"
        >WIN</v-chip
      >
      <v-chip v-else color="error" variant="elevated">LOSE</v-chip>
    </v-col>

    <v-col v-if="props.isDealer">Dealer</v-col>
    <v-col v-else>{{ (props.player as ResultPlayer).name }}</v-col>

    <v-col>{{ props.player.score }}</v-col>

    <v-col>
      <v-container>
        <v-row>
          <v-col v-for="card in props.player.hand">
            <Card :card="card" :size="4" />
          </v-col>
        </v-row>
      </v-container>
    </v-col>
  </v-row>
</template>

<style scoped lang="scss"></style>
