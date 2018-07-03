import Vue from 'vue'
import Vuex from 'vuex'

import pathify from './pathify'

Vue.use(Vuex)

import { authModule, authPluginMaker } from '@/auth'

const store = new Vuex.Store({
	modules: {
		auth: authModule
	},

	// state: {
	// },

	plugins: [pathify.plugin],
})
export default store

import router from './router'
authPluginMaker(router)(store)
