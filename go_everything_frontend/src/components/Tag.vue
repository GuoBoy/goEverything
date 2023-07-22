<template>
  <div class="tag">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">

</script>

<style lang="scss" scoped>
.tag {
  font: unset;
  width: fit-content;
  min-width: 40px;
  display: block;
  padding: 2px 5px;
  background: #12be8d;
  text-align: center;
}
</style>
