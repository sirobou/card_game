<script setup lang="ts">
import { defineProps } from "vue"

type Props = {
  size: number
}

const props = defineProps<Props>()
const setSize = (size: number) => `${size}rem`
</script>

<template>
  <div>
    <img src="joker_black.svg" />
  </div>
</template>

<style scoped lang="scss">
img {
  width: v-bind("setSize(props.size)");
}
</style>
