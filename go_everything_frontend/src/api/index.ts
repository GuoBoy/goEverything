import { CancelTokenSource } from "axios";
import { instance } from "./request"
import { Base64 } from 'js-base64';

const controller = new AbortController();

export const apiSearch = (key: string, cancelToken: CancelTokenSource) => {
  return new Promise((resolve, reject) => {
    instance.get(`/search?key=${Base64.encode(key)}`, { cancelToken: cancelToken.token }).then(res => resolve(res.data)).catch(e => reject(e))
  })
}

export const apiIndexDisk = (key: string) => {
  return new Promise((resolve, reject) => {
    instance.get(`/indexDisk?name=${Base64.encode(key)}`, { signal: controller.signal }).then(res => resolve(res.data)).catch(e => reject(e))
  })
}

export const apiGetDisks = () => {
  return new Promise((resolve, reject) => {
    instance.get("/disk", { signal: controller.signal }).then(res => resolve(res.data)).catch(e => reject(e))
  })
}

export const apiOpenPath = (path: string) => {
  return new Promise((resolve, reject) => {
    instance.get(`/open?path=${Base64.encode(path)}`, { signal: controller.signal }).then(res => resolve(res.data)).catch(e => reject(e))
  })
}
