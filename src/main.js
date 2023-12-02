import Vue from 'vue'
import VueMeta from 'vue-meta'
import App from './App.vue'
import router from './router'

Vue.use(VueMeta, {
  refreshOnceOnNavigation: true
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
