import Vue from 'vue'
import VueRouter from 'vue-router'
import TaskList from '@/components/TaskList.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'TaskList',
    component: TaskList
  }
]

const router = new VueRouter({
  routes
})

export default router
