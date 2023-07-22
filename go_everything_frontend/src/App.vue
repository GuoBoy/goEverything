<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter()
const toPath = (p: string) => router.push(p)

</script>

<template>
  <nav>
    <button @click="toPath('/')">文件搜索</button>
    <button @click="toPath('/diskManage')">硬盘索引</button>
  </nav>
  <router-view></router-view>
</template>
