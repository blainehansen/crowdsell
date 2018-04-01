import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false

import axios from 'axios'
import config from './config'

axios.defaults.baseURL = config.baseURL
axios.defaults.responseType = 'json'
// axios.interceptors.response.use(null, function (error) {
// 	return Promise.reject(error)
// })

import AsyncProperties from 'vue-async-properties'
Vue.use(AsyncProperties, {
	debounce: 1000
})

import store from '@/vuex'

import Cookies from 'js-cookie'
const authenticatedUser = Cookies.getJSON('authenticatedUser')
if (authenticatedUser) {
	authenticatedUser.userObj = { name: authenticatedUser.name }
	store.commit('loginUser', authenticatedUser)
}

new Vue({
	el: '#app',
	router,
	store,
	render: h => h(App)
})
