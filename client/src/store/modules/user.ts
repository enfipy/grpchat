const moduleState = {
  username: '',
}

const mutations = {
  setUsername(state: any, value: string) {
    state.username = value
  },
}

const getters = {
  getUsername: (state: any) => state.username,
}

export default {
  namespaced: true,
  state: moduleState,
  mutations,
  getters,
}
