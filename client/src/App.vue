<template>
  <div id="app">
    <chat />
    <modal-screen v-show="$store.state.user.username === ''" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

import Chat from '@/elements/Chat.vue'
import ModalScreen from '@/elements/ModalScreen.vue'

export default Vue.extend({
  name: 'app',
  components: {
    Chat,
    ModalScreen,
  },
})
</script>

<style lang="stylus">
body, html
  margin: 0
  padding: 0

#app
  font-family 'Avenir', Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
  text-align center
  background-color #ffffff
  color #2c3e50
  min-height 100vh
  display flex
  justify-content center
  align-items center
</style>
