import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: '/',
  routes: [
    { path: '/', componet: () => import('@/App') },
  ]
})


export default router