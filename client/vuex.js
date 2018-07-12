import Vue from 'vue'
import Vuex from 'vuex'

import pathify from '@/packages/vuex-pathify'

Vue.use(Vuex)

import { authModule, authPluginMaker } from '@/auth'
import projectModule from '@/projects/create/projectVuex'

const store = new Vuex.Store({
	modules: {
		auth: authModule,
		project: projectModule,
	},

	plugins: [pathify.plugin],
})
export default store

import router from './router'
authPluginMaker(router)(store)
