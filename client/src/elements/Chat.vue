<template>
  <div id="chat">
    <div id="messages">
      <message
        v-for="(msg, index) in $store.getters['chat/getMessages']"
        :key="index"
        :msg="msg"
      />
    </div>
    <form id="inputs" @submit.prevent="sendMessage">
      <input
        type="text"
        placeholder="Your message"
        v-model="message"
      />
      <input id="sendMessage" type="submit" value="Send" />
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapState } from 'vuex'

import config from '@/config'
import Message from '@/elements/Message.vue'

import { ChatPromiseClient } from 'schema/chat_grpc_web_pb'
import { ClientMessage, ServerMessage, GetMessagesRequest } from 'schema/chat_pb'

const client = new ChatPromiseClient(config.hostname, null, null)

export default Vue.extend({
  name: 'chat',
  components: {
    Message,
  },
  data: () => {
    return {
      message: 'first message',
    }
  },
  mounted() {
    const username = this.$store.state.user.username

    if (username !== '') {
      this.getMessages(username)
    } else {
      this.setupWatch(username)
    }
  },
  methods: {
    setupWatch(username: string) {
      this.$store.watch(
        (state: any, getters: any) => getters['user/getUsername'],
        (newValue, oldValue) => {
          console.log(oldValue)
          if (newValue !== '') {
            this.getMessages(newValue)
          }
        },
      )
    },
    getMessages(username: string) {
      if (username === '') {
        return
      }

      const req = new GetMessagesRequest()
      req.setUsername(username)
      const stream = client.getMessages(req, {})

      stream.on('data', (res: ServerMessage) => {
        this.$store.commit('chat/addMessage', res.toObject())
      })
      stream.on('error', (err: any) => {
        alert(`Failed to get messages: ${err.message}`)
      })
    },
    sendMessage() {
      if (this.message === '') {
        return
      }
      const username = this.$store.state.user.username
      const req = new ClientMessage()
      req.setMessage(this.message)

      client.sendMessage(req, { username }).catch((err: any) => {
        alert(`Failed to send message: ${err.message}`)
      })
    },
  },
})
</script>

<style scoped lang="stylus">
#chat
  box-shadow inset 0 0 5px #2c3e50
  background-color #fafafa
  border-radius 10px
  padding 20px
  min-height 600px
  min-width 400px
  display flex
  flex-direction column

#messages
  display flex
  flex-direction column

  .message
    list-style-type none

#inputs
  margin-top auto
</style>
