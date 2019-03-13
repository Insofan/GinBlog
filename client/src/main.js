// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import 'normalize.css'
import iView from 'iview'
import 'iview/dist/styles/iview.css'

import Vuetify from 'vuetify/lib'
import 'vuetify/dist/vuetify.min.css' // Ensure you are using css-loader

Vue.use(Vuetify, {
	theme: {
		primary: '#2196F3',
		secondary: '#8BC34A',
		accent: '#212121',
		error: '#FF5252',
		info: '#2196f3',
		success: '#4CAF50',
		warning: '#FFC107'
	}
})

Vue.use(iView)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
