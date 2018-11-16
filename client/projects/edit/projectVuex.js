import { makeMutations, makeGetters, genericSaveAction } from '@/vuexUtils'

const touchedKeyManifest = []

const projectState = {
	name: null,
	description: null,
	story: null,
	promises: [],
	category: null,
}

const state = {
	id: null,
	...projectState
}

import { secureApi } from '@/api'

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
			const response = await secureApi.saveProject(state.id, projectPatches)

			if (response.data) {
				const projectId = response.data

				commit('SET_ID', projectId)
				return projectId
			}
		}),
	}
}
