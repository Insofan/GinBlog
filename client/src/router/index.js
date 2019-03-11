import Vue from 'vue'
import Router from 'vue-router'
import HomeController from '@/components/Home/HomeController'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HomeController',
      component: HomeController
    }
  ]
})
