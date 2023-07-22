<template>
  <div class="disk">
    <div class="name">
      {{ data.name }}
    </div>
    <button @click="onIndex(data.name)">重新索引</button>
    <div class="description">
      <div class="item">文件: <tag>{{ data.file_count }}</tag>
      </div>
      <div class="item">文件夹:<tag>{{ data.dir_count }}</tag>
      </div>
    </div>
    <div class="updated_at">
      {{ dayjs(data.updated_at).format("YYYY-MM-DD hh:mm:ss") }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { apiIndexDisk } from '../api';
import dayjs from "dayjs"

defineProps<{
  data: DiskModel
}>()

const onIndex = (name: string) => {
  apiIndexDisk(name).then((res: any) => {
    console.log(res);
  })
}
</script>

<style lang="scss" scoped>
.disk {
  margin: 10px;
  padding: 10px;
  outline: 1px solid #0606;
  border-radius: 10px;

  .name {
    font-weight: bolder;
  }

  .description {
    font-weight: bolder;
    display: flex;

    .item {
      display: flex;
      margin: 10px;
    }
  }
}
</style>
