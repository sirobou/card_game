<script setup lang="ts">
import Card from "@/components/Card.vue"
import CardBack from "@/components/CardBack.vue"
import { Dealer } from "@/types/Player"
import { CardColor } from "@/utils/Color"

type Props = {
  dealer: Dealer
  is1on1: boolean
}

const props = defineProps<Props>()
</script>

<template>
  <v-card :color="CardColor(dealer)">
    <template v-slot:title>
      <div :class="props.is1on1 ? 'name-1on1' : ''">
        <v-chip v-if="props.dealer.isStand" color="primary" variant="elevated"
          >STAND</v-chip
        >
        <span>Dealer</span>
      </div>
    </template>
    <div
      :class="props.is1on1 ? 'dealer-hand dealer-hand--1on1' : 'dealer-hand'"
    >
      <Card :card="props.dealer.publicHand" :size="9.5" />
      <CardBack v-for="_ in new Array(dealer.handCount - 1)" :size="9.5" />
    </div>
  </v-card>
</template>

<style scoped lang="scss">
.v-chip {
  margin-right: 1.5rem;
}
.v-card {
  height: 100%;
  padding-top: 1rem;
  padding-left: 1rem;
  padding-right: 0.5rem;
  padding-bottom: 1rem;
}

.dealer-hand {
  display: flex;
  flex-wrap: wrap;
  *:nth-child(n) {
    margin-right: 0.5rem;
  }

  &--1on1 {
    justify-content: center;
  }
}

.name-1on1 {
  text-align: center;
}
</style>
