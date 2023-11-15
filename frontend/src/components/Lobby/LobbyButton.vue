<script setup lang="ts">
type Emit = {
  click: []
}

const emit = defineEmits<Emit>()
const onClick = () => {
  emit("click")
}
</script>

<template>
  <v-btn size="x-large" color="primary" @click="onClick">ゲーム開始</v-btn>
</template>
