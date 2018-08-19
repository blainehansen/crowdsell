import { makeMutations, makeGetters, genericSaveAction } from '@/vuexUtils'
import { has } from 'lodash'

const touchedKeyManifest = []

const projectState = {
	name: null,
	description: null,
}

const state = {
	id: null,
	...projectState
}

import { publicApi } from '@/api'

export default {
	namespaced: true,

	state,

	mutations: {
		SET_ID(state, id) {
			state.id = id
		},
		...makeMutations(state, Object.keys(projectState), touchedKeyManifest),
	},

	getters: {
		...makeGetters(touchedKeyManifest),
	},

	actions: {
		saveProject: genericSaveAction(touchedKeyManifest, async function({ state, getters, commit }, projectPatches) {
			const response = await publicApi.saveProject(state.id, projectPatches)

			if (has(response, 'data.id'))
				commit('SET_ID', response.data.id)
		}),
	}
}
