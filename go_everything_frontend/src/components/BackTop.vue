<template>
  <div class="backtop" @click="backTop"> ^ </div>
</template>

<script setup lang="ts">
const props = defineProps<{ el: any }>()

const backTop = () => {
  props.el.scrollTo({ top: 0, behavior: "smooth" })
}
</script>

<style lang="scss">
.backtop {
  position: absolute;
  right: 40px;
  bottom: 40px;
  height: 40px;
  width: 40px;
  font-size: xx-large;
  font-weight: 900;
  text-align: center;
  cursor: pointer;
  z-index: 9999;
}

.backtop:hover {
  background: #9f9c9c;
  border-radius: 50%;
}
</style>
