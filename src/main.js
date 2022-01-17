import Vue from 'vue'
import VueMeta from 'vue-meta'
import App from './App.vue'

Vue.use(VueMeta, {
  refreshOnceOnNavigation: true
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
