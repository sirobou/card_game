<script setup lang="ts">
import Card from "@/components/Card.vue"
import { Card as Card_t } from "@/types/Card"
type Props = {
  hand: Card_t[]
}

const props = defineProps<Props>()
</script>

<template>
  <div class="hand">
    <span v-for="card in props.hand">
      <Card :card="card" :size="9.5" />
    </span>
  </div>
</template>

<style scoped lang="scss">
.hand {
  display: flex;
  flex-wrap: wrap;
  *:nth-child(n) {
    margin-left: 0.2rem;
  }
}
</style>
