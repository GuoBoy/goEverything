import { ref } from "vue"
import { apiSearch } from "../api"
import { CancelTokenSource } from "axios"

interface HistoryItemT {
  key: string
  value: SearchResultModel
}

export const useSearchCache = () => {
  const histories = ref<HistoryItemT[]>([])

  const push = (item: HistoryItemT) => {
    for (let i = 0; i < histories.value.length; i++) {
      const itm = histories.value[i].key
      if (itm === item.key) return
    }
    histories.value.push(item)
  }

  const search = (key: string, cancelToken: CancelTokenSource) => {
    return new Promise((resolve) => {
      for (let i = 0; i < histories.value.length; i++) {
        const itm = histories.value[i]
        if (itm.key === key) {
          resolve(itm.value)
          return
        }
      }
      apiSearch(key, cancelToken).then(((res: any) => {
        push({ key: key, value: res.data })
        resolve(res.data)
      }))
    })

  }

  const clear = () => {
    histories.value = []
  }

  return {
    search, clear
  }
}