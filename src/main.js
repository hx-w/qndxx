import Vue from 'vue'
import VueMeta from 'vue-meta'
import App from './App.vue'
import { VueJsonp } from 'vue-jsonp'

Vue.use(VueMeta, {
  refreshOnceOnNavigation: true
})
Vue.use(VueJsonp)

Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
