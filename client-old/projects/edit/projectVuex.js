import { makeMutations, makeGetters, genericSaveAction } from '@/vuexUtils'

const touchedKeyManifest = []

const projectState = {
	name: null,
	description: null,
	story: null,
	promises: [],
	category: null,
	managedImages: {},
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
		saveProject: genericSaveAction(touchedKeyManifest, async function({ state, commit }, projectPatches) {
			const managedImages = projectPatches.managedImages
			if (managedImages) {
				const projectId = state.id
				if (!projectId) throw new Error("blaine, you've let an uncreated project get created with managedImages")

				const finalManagedImages = {}

				for (const key in managedImages) {
					const managedImage = managedImages[key]
					const newUrl = secureApi.uploadProjectImage(projectId, managedImage.hash, managedImage.file)
					finalManagedImages[key] = { ...managedImage, url: newUrl, uploaded: true }
				}

				projectPatches.managedImages = finalManagedImages
				commit('SET_MANAGED_IMAGES', finalManagedImages)
			}

			const response = await secureApi.saveProject(state.id, projectPatches)

			if (response.data) {
				const projectId = response.data

				commit('SET_ID', projectId)
				return projectId
			}
		}),
	}
}
