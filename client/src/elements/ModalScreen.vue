<template>
  <transition name="modal">
    <div id="modal-screen">
        <div class="modal-container" @click.stop>
          <form @submit.prevent="setUsername">
            <input
              type="text"
              placeholder="Username"
              v-model="username"
            />
            <input id="setUsername" type="submit" value="Enter grpchat" />
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
    min-width: 320px
    min-height: 240px

    background: #2c3e50
    border-radius: 5px
    box-shadow: 0 3px 8px rgba(0, 0, 0, .33)
    transition: all .2s ease
</style>
