/// <reference types="vite/client" />
declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const vueComponent: DefineComponent<{}, {}, any>;
  export default vueComponent;
}

interface DiskModel {
  id: string
  name: string
  updated_at: string
  file_count: number
  dir_count: number
}

interface SearchResultModel {
  id: string
  path: string
  name: string
  size: number
  type: number
}