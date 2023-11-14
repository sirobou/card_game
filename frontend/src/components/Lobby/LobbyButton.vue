<script setup lang="ts">
import { defineEmits } from "vue"

type Emit = {
  push: []
}

const emit = defineEmits<Emit>()
const onPush = () => {
  emit("push")
}
</script>

<template>
  <v-btn size="x-large" color="primary" @push="onPush">ゲーム開始</v-btn>
</template>
