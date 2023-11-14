<script setup lang="ts">
import Title from "@/components/Title.vue"
import HomeForm from "@/components/Home/HomeForm.vue"
import { BASE_URL } from "@/consts/const"
import router from "@/router"

const submit = async (name: string) => {
  const resp = await fetch(`${BASE_URL}/api/lobby/join`, {
    method: "POST",
    body: JSON.stringify({ name }),
  })

  if (resp.ok) {
    const { id } = await resp.json()
    localStorage.setItem("id", id)
    router.push("/lobby")
  } else {
    console.error(resp)
    alert("エラーが発生しました")
  }
}
</script>

<template>
  <Title />
  <HomeForm @submit="submit" />
</template>
