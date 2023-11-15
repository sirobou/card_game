<script setup lang="ts">
import SVGElement from "@/components/SVGElement.vue"
type Props = {
  size: number
}

const props = defineProps<Props>()
const setSize = (size: number) => `${size}rem`
</script>

<template>
  <div>
    <SVGElement name="card_back" />
  </div>
</template>

<style scoped lang="scss">
img {
  width: v-bind("setSize(props.size)");
}
</style>
