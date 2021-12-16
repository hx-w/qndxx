import Vue from 'vue'
import App from './App.vue'
import VueMeta from 'vue-meta'

Vue.use(VueMeta, {
  // optional pluginOptions
  refreshOnceOnNavigation: true
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
