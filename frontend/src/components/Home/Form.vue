<script setup lang="ts">
import { ref, computed } from "vue"
const name = ref("")
const rules = [
  (v: string) => {
    if (!v) return "プレイヤー名を入力してください"
    return true
  },
  (v: string) => {
    if (v.length >= 16) return "プレイヤー名は16文字以内で入力してください"
    return true
  },
]

const isButtonDisabled = computed(() =>
  rules.some((rule) => rule(name.value) !== true)
)

const submit = () => {
  alert(`プレイヤー名: ${name.value}`)
}
</script>

<template>
  <v-form fast-fail @submit.prevent="submit">
    <v-container>
      <v-row align="center" justify="center">
        <v-col>
          <v-text-field
            v-model="name"
            :rules="rules"
            as="input"
            type="text"
            label="プレイヤー名"
          />
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-btn
          type="submit"
          :disabled="isButtonDisabled"
          size="x-large"
          color="primary"
          >ゲームに参加</v-btn
        >
      </v-row>
    </v-container>
  </v-form>
</template>

<style scoped lang="scss">
.v-col {
  max-width: 40rem;
  min-width: 20rem;
}
</style>
