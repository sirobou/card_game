<script setup lang="ts">
type Props = {
  icon: string
  name: string
}
const { icon, name } = defineProps<Props>()
</script>

<template>
  <v-card variant="outlined">
    <template v-slot:title>
      <v-icon size="x-large" color="secondary">{{ icon }}</v-icon>
      <span class="name-text">{{ name }}</span>
    </template>
  </v-card>
</template>

<style scoped lang="scss">
.name-text {
  margin-left: 2rem;
  font-size: 1.5rem;
}
</style>
