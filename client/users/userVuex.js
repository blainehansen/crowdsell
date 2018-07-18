import { privateApi } from '@/api'

const userState = {
	name: '',
	bio: '',
	links: [],
	location: '',
}

const state = {
	$fetched: false,
	...userState,
}

export default {
	namespaced: true,

	state,

	mutations: {
		SET_FULL_USER(state, fullUser) {
			state.name = fullUser.name

			state.$fetched = true
		},

		...makeMutations(state, Object.keys(projectState)),
	},

	actions: {
		async fetchInitial({ state, commit, getters }) {
			// TODO perhaps this should be stored in localstorage?
			if (state.$fetched) return

			const { data: fullUser } = await privateApi.fetchFullUser(getters['auth/userSlug'])

			commit('SET_FULL_USER', fullUser)
		}
	}
}