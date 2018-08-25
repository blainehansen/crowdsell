import Vue from 'vue'
import App from './App'

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

import Cookies from 'js-cookie'
const signedUser = Cookies.getJSON('signedUser')
if (signedUser) {
	store.commit('auth/login', signedUser)
}

import { formatSpacesUrl } from './utils'
Vue.filter('formatSpacesUrl', formatSpacesUrl)

new Vue({
	el: '#app',
	store,
	router,
	render: h => h(App)
})


// https://forum.vuejs.org/t/how-import-a-cdn-on-vuejs/6824/2
