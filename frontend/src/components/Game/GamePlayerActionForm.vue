<script setup lang="ts">
import { defineEmits, defineProps } from "vue"

const props = defineProps<{
  disabled: boolean
}>()

type Emit = {
  hit: []
  stand: []
}
const emit = defineEmits<Emit>()

const hit = () => emit("hit")
const stand = () => emit("stand")
</script>

<template>
  <div>
    <v-btn color="primary" @click="hit" :disabled="props.disabled">Hit</v-btn>
    <v-btn color="primary" @click="stand" :disabled="props.disabled"
      >Stand</v-btn
    >
  </div>
</template>

<style scoped lang="scss">
.v-btn {
  margin-right: 0.5rem;
  width: 5rem;
}
</style>
