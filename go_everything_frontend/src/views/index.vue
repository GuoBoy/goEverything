<template>
  <div class="container">
    <input type="text" style="width: 50vw;padding: 2vh;" v-model="searchKey" @input="onSearch" placeholder="输入文件/目录关键词">
    <div class="table">
      <div class="header row">
        <div>Name</div>
        <div>Path</div>
        <div>Size</div>
        <div>Open</div>
      </div>
      <ul class="body" ref="backTopEl">
        <li class="row" v-for="item in items" :key="item.id">
          <div>
            <div v-html="highLight(item.name)"></div>
          </div>
          <div>
            <div v-html="highLight(item.path)"></div>
          </div>
          <div>{{ item.size }}B</div>
          <div>
            <button @click="apiOpenPath(getFilePath(item))">DIR</button><button v-if="!item.type"
              @click="apiOpenPath(item.path)">FILE</button>
          </div>
        </li>
      </ul>
    </div>
  </div>
  <BackTop :el="backTopEl"></BackTop>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useSearchCache } from "./hooks";
import { apiOpenPath } from "../api"
import axios, { CancelTokenSource } from "axios";
import BackTop from "../components/BackTop.vue"

const searchKey = ref("")
const searchCache = useSearchCache()
const items = ref<SearchResultModel[]>([])

const cancelToken = ref<CancelTokenSource>()

const onSearch = () => {
  if (!searchKey.value) return
  cancelToken.value?.cancel()
  cancelToken.value = axios.CancelToken.source()
  searchCache.search(searchKey.value, cancelToken.value).then((res: any) => items.value = res)
}

const getFilePath = (item: SearchResultModel): string => {
  if (item.type) return item.path
  return item.path.substring(0, item.path.lastIndexOf('\\'))
}

const highLight = (text: string) => {
  if (searchKey.value === "") return text
  const regx = new RegExp(searchKey.value, "ig")
  return text.replace(regx, `<span style="background:yellow;">${searchKey.value}</span>`)
}

const backTopEl = ref(null)

</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;

  .table {
    display: grid;
    grid-template-rows: auto 1fr;
    overflow-y: auto;

    .header {
      background-color: #409eff;
    }

    .row {
      $name: 200px;
      $size: 100px;
      $open: 200px;
      $margin: 120px;
      display: grid;
      grid-template-columns: $name calc(100vw - $name - $size - $open - $margin) $size $open;
      word-wrap: break-word;
      border-top: 1px solid #409edd;
      padding: 10px 0;

      div:nth-child(2n+1) {
        border-right: 1px solid #409edd;
      }
    }

    .body {
      max-height: 78vh;
      overflow-y: auto;
      list-style: none;
      margin: 0;
      padding: 0;
    }
  }
}
</style>