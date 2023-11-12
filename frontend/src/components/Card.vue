<script setup lang="ts">
import { defineProps } from "vue"
import { Card } from "@/types/Card"
import { convertCardImage } from "@/utils/Card"

type Props = {
  card: Card
  size: number
}

const props = defineProps<Props>()

const cardImage = convertCardImage(props.card)

const setSize = (size: number) => `${size}rem`
</script>

<template>
  <div>
    <img :src="cardImage" />
  </div>
</template>

<style scoped lang="scss">
img {
  width: v-bind("setSize(props.size)");
}
</style>
