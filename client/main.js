import Vue from 'vue'
import App from './App'
import router from './router'

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false

import AsyncProperties from 'vue-async-properties'
Vue.use(AsyncProperties, {
	debounce: 1000
})

import store from '@/vuex'

import Cookies from 'js-cookie'
const authToken = Cookies.getJSON('authToken')
if (authToken) {
	store.commit('login', authToken)
}

new Vue({
	el: '#app',
	router,
	store,
	render: h => h(App)
})
