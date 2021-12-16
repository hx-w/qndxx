import Vue from 'vue'
import App from './App.vue'
import VueMeta from 'vue-meta'
import axios from 'axios'

Vue.use(VueMeta, {
  refreshOnceOnNavigation: true
})

Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
