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
  margin 5px
  padding 5px
  border-radius 5px

  .username
    color #c1c3ce
  .body
    color #000000
  .time
    color #777777
</style>
