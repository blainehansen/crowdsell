import Vue from 'vue'
import App from './App'
import './registerServiceWorker'

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false

import AsyncProperties from 'vue-async-properties'
Vue.use(AsyncProperties, {
	debounce: 1000
})

import store from './vuex'
import router from './router'

// store.commit('auth/login', {
// 	token: 'eyJpIjoiWk5XR292UG4iLCJlIjoxNTM5Mzc0Njc1fQ.QfrTb6QH1wYAVbpp5PS5WtPk-G4VbGjDcM449JbM1AQ',
// 	name: 'dude',
// 	email: 'dude@gmail.com',
// 	slug: 'dude',
// })

import Cookies from 'js-cookie'
const signedUser = Cookies.getJSON('signedUser')
if (signedUser) {
	store.commit('auth/login', signedUser)
}

import { formatProfileImageUrl } from './utils'
Vue.filter('formatProfileImageUrl', formatProfileImageUrl)


new Vue({
	router,
	store,
	render: h => h(App)
}).$mount('#app')
