<template>
  <input type="text" style="width: 50vw;padding: 2vh;" v-model="directory" placeholder="输入文件夹绝对路径(eg.C:/)"><button
    @click="onIndex">开始索引</button>
  <div class="disk-container">
    <Disk v-for="itm in disks" :data="itm" :key="itm.id"></Disk>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import Disk from "../components/Disk.vue"
import { apiIndexDisk, apiGetDisks } from "../api";

const disks = ref<DiskModel[]>([])
const refreshDisk = () => {
  apiGetDisks().then((res: any) => {
    disks.value = res.data
  })
}

const directory = ref("")
const onIndex = () => {
  if (!directory.value) return
  apiIndexDisk(directory.value).then(res => {
    console.log(res);
  })
}

onMounted(() => {
  refreshDisk()
})

</script>

<style lang="scss" scoped>
.disk-container {
  display: flex;
  flex-wrap: wrap;
}
</style>