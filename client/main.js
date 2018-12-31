import Vue from 'vue'
import App from './App'
import './registerServiceWorker'
Vue.config.productionTip = false

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)
import 'bootstrap-vue/dist/bootstrap-vue.css'


import { library } from '@fortawesome/fontawesome-svg-core'
import { faDove } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faDove)

Vue.component('fa-icon', FontAwesomeIcon)

import SignUpForm from '@/components/SignUpForm'
Vue.component('SignUpForm', FontAwesomeIcon)

import OutLink from '@/components/OutLink'
Vue.component('out-link', OutLink)

import router from './router'

// import Cookies from 'js-cookie'
// const signedUser = Cookies.getJSON('signedUser')
// if (signedUser) {
// 	store.commit('auth/login', signedUser)
// }


Vue.directive('visible', (el, binding) => {
	el.style.visibility = binding.value ? 'visible' : 'hidden'
})

Vue.directive('displaying', (el, binding) => {
	el.style.display = binding.value ? '__invalid' : 'none'
})


new Vue({
	router,
	render: h => h(App)
}).$mount('#app')
