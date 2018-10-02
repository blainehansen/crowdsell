import { Payload } from 'vuex-pathify'
import constantCase from 'constant-case'

// TODO this doesn't recursively account for nested objects
export function makeMutations(stateObject, relevantKeys, touchedKeyManifest) {
	const mutations = {}

	for (let i = relevantKeys.length - 1; i >= 0; i--) {
		const key = relevantKeys[i]
		const touchedKey = `${key}$touched`
		stateObject[touchedKey] = false

		const mutationName = 'SET_' + constantCase(key)
		mutations[mutationName] = function(state, value) {
      state[key] = value instanceof Payload
        ? value.update(state[key])
        : value
      state[touchedKey] = true
      state.$touchedCount++
    }

    // mutations[mutationName + '_RESET'] = function(state) {
    //   state[touchedKey] = false
    //   state.$touchedCount--
    // }

    touchedKeyManifest.push(touchedKey)

		stateObject.$touchedCount = 0

		mutations.RESET = function(state) {
			for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
				state[touchedKeyManifest[i]] = false
			}
			state.$touchedCount = 0
		}
	}

	return mutations
}

export function makeGetters(touchedKeyManifest) {
	return {
		$anyTouched: (state) => state.$touchedCount > 0,
	}
}

export function genericSaveAction(touchedKeyManifest, saveFunction) {
	return async function(context) {
		const { state, getters, commit } = context

		if (!getters.$anyTouched) return

		const patches = {}
		for (let i = touchedKeyManifest.length - 1; i >= 0; i--) {
			const touchedKey = touchedKeyManifest[i]
			if (state[touchedKey]) {
				const actualKey = touchedKey.replace(/\$touched$/, '')
				patches[actualKey] = state[actualKey]
			}
		}

		const saveResult = await saveFunction(context, patches)

		commit('RESET')

		return saveResult
	}
}
