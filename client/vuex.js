import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import { authModule, authPlugin } from '@/auth'

export default new Vuex.Store({
	modules: {
		auth: authModule
	},

	// state: {
	// },

	plugins: [authPlugin],
})
