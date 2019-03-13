import Vue from 'vue'
import Router from 'vue-router'
import HomeController from '@/components/Front/Home/HomeController'
import LoginController from '@/components/Front/Login/LoginController'
import AdminMainController from '@/components/Admin/Main/AdminMainController'

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
    },
    {
      path: '/admin',
      name: 'AdminMainController',
      component: AdminMainController
    }
  ]
})
