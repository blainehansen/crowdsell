import Vue from 'vue'

const components = require.context('@/components', false, /\w+\.(vue)$/)

for (const fileName of components.keys()) {
	const componentConfig = components(fileName)
	const componentName = fileName.split('/').pop().split('.')[0]
	Vue.component(componentName, componentConfig.default || componentConfig)
}


Vue.directive('visible', (el, binding) => {
	el.style.visibility = binding.value ? 'visible' : 'hidden'
})

Vue.directive('displaying', (el, binding) => {
	el.style.display = binding.value ? '__invalid' : 'none'
})



// import Cookies from 'js-cookie'
// const signedUser = Cookies.getJSON('signedUser')
// if (signedUser) {
// 	store.commit('auth/login', signedUser)
// }
