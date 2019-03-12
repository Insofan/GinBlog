import Vue from 'vue'
import Router from 'vue-router'
import HomeController from '@/components/Front/Home/HomeController'
import LoginController from '@/components/Front/Login/LoginController'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HomeController',
      component: HomeController
    },
    {
      path: '/login',
      name: 'LoginController',
      component: LoginController
    }
  ]
})
