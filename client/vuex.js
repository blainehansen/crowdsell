import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import { authModule, authPlugin } from '@/router'

export default new Vuex.Store({
	modules: {
		auth: authModule
	},

	// state: {
	// },

	plugins: [authPlugin],
})
