<template>
  <div id="app">
    <div id="chat">
      <div id="messages">
        <ul>
          <li class="message" v-for="(msg, index) in messages" :key="index">
            {{ msg.message }}. By {{ msg.username }}. At {{ msg.createdAt }}
          </li>
        </ul>
      </div>
      <form id="inputs" @submit.prevent="sendMessage">
        <input
          type="text"
          placeholder="Your message"
          v-model="message"
        />
        <input id="sendMessage" type="submit" value="Send" />
      </form>
      <form @submit.prevent="getMessages">
        <input
          type="text"
          placeholder="Username"
          v-model="username"
        />
        <input id="getMessages" type="submit" value="Get messages" />
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import grpcWeb from 'grpc-web'

import { ChatPromiseClient } from 'schema/chat_grpc_web_pb'
import { ClientMessage, ServerMessage, GetMessagesRequest, SendMessageResponse } from 'schema/chat_pb'

const client = new ChatPromiseClient('http://localhost:50051', null, null)

export default Vue.extend({
  name: 'chat',
  data: () => {
    const messages: any[] = []
    return {
      messages,
      message: 'first message',
      username: 'doe',
    }
  },
  methods: {
    getMessages() {
      if (this.username === '') {
        return
      }
      this.messages = []

      const req = new GetMessagesRequest()
      req.setUsername(this.username)

      const stream = client.getMessages(req, {})
      stream.on('data', (res: ServerMessage) => {
        console.log(res.getUsername())
        this.messages.push(res.toObject())
      })
      stream.on('error', (err: any) => {
        alert(err.message)
      })
    },
    sendMessage() {
      if (this.message === '') {
        return
      }

      const req = new ClientMessage()
      req.setMessage(this.message)
      const metadata = {
        username: this.username,
      }
      client.sendMessage(req, metadata).catch((err: any) => {
        console.log(err)
      })
    },
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

#chat
  background-color #fafafa
  border-radius 10px
  padding 20px
  min-height 400px
  display flex
  flex-direction column

#messages
  display flex

  .message
    list-style-type none

#inputs
  margin-top auto
</style>
