import axios from "axios"

export const instance = axios.create({
  baseURL: "http://localhost:23456"
})

instance.interceptors.request.use(config => {
  // config.headers ? config.headers['Authorization'] = userStore.token : ""
  return config
})
