<template>
  <component :is="svg" />
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
const props = defineProps<{ name: string }>()
const module = async () => await import(`@/assets/svg/${props.name}.svg`)
const svg = ref<typeof module | null>(null)
onMounted(async () => {
  svg.value = await module()
})
</script>
