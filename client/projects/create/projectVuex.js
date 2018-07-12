// import { isPlainObject } from '@/packages/vuex-pathify/utils/object'
import { resolveName } from '@/packages/vuex-pathify/services/resolver'
import Payload from '@/packages/vuex-pathify/classes/Payload'
import { hasValue } from '@/packages/vuex-pathify/utils/object'

// we want to create a system that allows us to specify state variables whose mutations will automatically set a "$touched"
const touchedKeyManifest = []

function makeMutations(stateObject, relevantKeys) {
	const mutations = {}

	for (let i = relevantKeys.length - 1; i >= 0; i--) {
		const key = relevantKeys[i]
		const touchedKey = `${key}$touched`
		stateObject[touchedKey] = false

		// const isNestedObject = isPlainObject(value)
		const mutationName = resolveName('mutations', key)
		mutations[mutationName] = function(state, value) {
      state[key] = value instanceof Payload
        ? value.update(state[key])
        : value
      state[touchedKey] = true
      // if (isNestedObject) stateObject[`${key}$anyTouched`] = false
    }

    mutations[mutationName + '_RESET'] = function(state) {
      state[touchedKey] = false
    }

    touchedKeyManifest.push(touchedKey)

		// if (isNestedObject) {
		// 	stateObject[`${key}$anyTouched`] = false
		// 	makeMutations(value, mutations)
		// }
	}

	return mutations
}


const projectState = {
	name: null,
	description: null,
}

const state = {
	id: null,
	...projectState
}

import api from '@/api'

export default {
	namespaced: true,

	state,

	mutations: {
		SET_ID(state, id) {
			state.id = id
		},
		...makeMutations(state, Object.keys(projectState)),

		RESET(state) {
			for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
				state[touchedKeyManifest[i]] = false
			}
		}
	},

	getters: {
		$anyTouched(state) {
			for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
				if (state[touchedKeyManifest[i]]) return true
			}
			return false
		}
	},

	actions: {
		async saveProject({ state, getters, commit }) {
			if (!getters.$anyTouched) return

			const projectPatches = {}
			for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
				const touchedKey = touchedKeyManifest[i]
				if (state[touchedKey]) {
					const actualKey = touchedKey.replace(/\$touched$/, '')
					projectPatches[actualKey] = state[actualKey]
				}
			}

			const response = await api.saveProject(state.id, projectPatches)
			if (hasValue(response, 'data.id')) commit('SET_ID', response.data.id)

			commit('RESET')
		}
	}
}
