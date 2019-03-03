import Vue from 'vue'
import VueChatScroll from 'vue-chat-scroll'

import store from '@/store'
import App from '@/App.vue'

Vue.use(VueChatScroll)
Vue.config.productionTip = false

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app')
