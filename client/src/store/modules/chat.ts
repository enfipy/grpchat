const messages: string[] = []

const moduleState = {
  messages,
}

const mutations = {
  addMessage(state: any, value: any) {
    if (state.messages.length >= 100) { state.messages.shift() }

    // Todo: Optimize this shit :/
    state.messages.push(value)
    state.messages.sort((a: any, b: any) => {
      if (a.createdAt < b.createdAt) {
        return -1
      }
      if (a.createdAt > b.createdAt) {
        return 1
      }
      return 0
    })
  },
}

const getters = {
  getMessages: (state: any) => state.messages,
}

export default {
  namespaced: true,
  state: moduleState,
  mutations,
  getters,
}
