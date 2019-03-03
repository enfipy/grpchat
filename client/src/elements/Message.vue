<template>
  <div class="message">
    <b class="username">{{ msg.username }}: </b>
    <span class="body">{{ msg.message }}</span>
    <span class="time">{{ msg.createdAt | formatUnix }}</span>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import moment from 'moment'

export default Vue.extend({
  props: ['msg'],
  filters: {
    formatUnix(timestamp: number): string {
      const date = moment.unix(timestamp).format('HH:mm')
      return date
    },
  },
})
</script>

<style scoped lang="stylus">
.message
  display flex
  justify-content space-between
  background-color #ffffff
  margin-bottom 10px
  padding 5px 10px
  border-radius 5px
  box-shadow inset 0 0 1px #999999

  .username
    color #8d8d8d
  .body
    max-width 300px
    color #3e3e3e
  .time
    color #8d8d8d
</style>
