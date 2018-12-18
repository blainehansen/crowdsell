import Vue from 'vue'
import App from './App'
import './registerServiceWorker'
Vue.config.productionTip = false

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)
import 'bootstrap-vue/dist/bootstrap-vue.css'

import vClickOutside from 'v-click-outside'
Vue.use(vClickOutside)

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

Vue.directive('visible', (el, binding) => {
	el.style.visibility = binding.value ? 'visible' : 'hidden'
})

Vue.directive('displaying', (el, binding) => {
	el.style.display = binding.value ? '__invalid' : 'none'
})


// TODO see if you can do a webpack multiple import thing
import SingleSelect from '@/components/SingleSelect'
import MultipleSelect from '@/components/MultipleSelect'

const components = [
	SingleSelect,
	MultipleSelect,
]

for (const component of components) {
	Vue.component(component.name, component)
}



new Vue({
	router,
	store,
	render: h => h(App)
}).$mount('#app')
