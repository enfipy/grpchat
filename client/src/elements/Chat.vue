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
        id="message_box"
        placeholder="Your message"
        v-model="message"
      />
      <input id="send_message" type="submit" value="Send" />
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

      this.message = ''
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
  max-height 600px
  min-width 400px
  display flex
  flex-direction column

#messages
  height 100%
  overflow-y scroll

#inputs
  min-height 40px
  display flex
  border-top 1px #41b883
  margin-top 20px

  #message_box
    margin-right 10px
    padding 0 10px

    border-radius 3px
    border none
    box-shadow inset 0 0 5px #999999
    width 70%
    font-size 16px

    &:focus
      outline: none
      caret-color: #41b883
      box-shadow inset 0 0 5px #41b883

  #send_message
    flex: 1
    padding: 10px
    border-radius 5px
    font-size 16px
</style>
