<script setup lang="ts">
type Props = {
  size: number
}

const props = defineProps<Props>()
const setSize = (size: number) => `${size}rem`
</script>

<template>
  <div>
    <img src="card_back.svg" />
  </div>
</template>

<style scoped lang="scss">
img {
  width: v-bind("setSize(props.size)");
}
</style>
