<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
})
</script>

<style scoped lang="stylus">
h3
  margin 40px 0 0
</style>
