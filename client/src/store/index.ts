import Vue from 'vue'
import Vuex from 'vuex'
import cookies from 'js-cookie'
import createPersistedState from 'vuex-persistedstate'

import moduleUser from './modules/user'
import moduleChat from './modules/chat'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    user: moduleUser,
    chat: moduleChat,
  },
  plugins: [
    createPersistedState({
      paths: ['user'],
      storage: {
        getItem: (key) => cookies.get(key),
        setItem: (key, value) => cookies.set(key, value, { expires: 3 }),
        removeItem: (key) => cookies.remove(key),
      },
    }),
  ],
})
