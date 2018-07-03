// import { isPlainObject } from 'vuex-pathify/src/utils/object'
import { resolveName } from 'vuex-pathify/src/services/resolver'
import Payload from 'vuex-pathify/src/classes/Payload'

// we want to create a system that allows us to specify state variables whose mutations will automatically set a "$touched"
const touchedKeyManifest = []

function makeMutations(stateObject, mutations = {}) {
	for (const [key, value] of Object.entries(stateObject)) {
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

const state = {
	name: null,
}

export default {
	namespaced: true,

	state,

	mutations: {
		...makeMutations(state),
		RESET(state) {
			for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
				state[touchedKeyManifest[i]] = false
			}
		}
	},
}
