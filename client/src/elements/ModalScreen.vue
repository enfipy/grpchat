<template>
  <transition name="modal">
    <div id="modal-screen">
        <div class="modal-container" @click.stop>
          <img id="logo" src="../assets/logo.svg" />
          <form id="inputs" @submit.prevent="setUsername">
            <input
              id="username"
              type="text"
              placeholder="Username"
              v-model="username"
            />
            <input id="set_username" type="submit" value="Enter chat" />
          </form>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data: () => {
    return {
      username: '',
    }
  },
  computed: {
    modalShow() {
      return this.$store.state.user.username
    },
  },
  methods: {
    setUsername() {
      if (this.username === '') {
        alert('Invalid username')
      }
      this.$store.commit('user/setUsername', this.username)
    },
  },
}
</script>

<style lang="stylus" scoped>
#modal-screen
  position: fixed
  width: 100%
  height: 100%

  z-index: 9998
  top: 0
  left: 0

  display: flex
  justify-content: center
  align-items: center

  background-color: rgba(0, 0, 0, .66)
  transition: opacity .2s ease

  &.modal-enter, &.modal-leave-to
    opacity: 0

  &.modal-enter .modal-container,
  &.modal-leave-to .modal-container
    transform: scale(1.05)

  .modal-container
    padding 30px 20px

    background: #ffffff
    border-radius: 5px
    box-shadow: 0 3px 8px rgba(0, 0, 0, .33)
    transition: all .2s ease

    #logo
      margin-bottom 30px

    #inputs
      min-height 40px
      display flex
      border-top 1px #41b883

      #username
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

      #set_username
        flex: 1
        padding: 10px
        border-radius 5px
        font-size 16px
</style>
