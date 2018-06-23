import Cookies from 'js-cookie'
import api, { privateHttp } from '@/api'
import router from '@/router'

function decodeToken(token) {
	const segments = token.split('.')
	if (segments.length != 2) console.error('invalid token', token)
	return JSON.parse(atob(segments[0]))
}

export const authModule = {
	state: {
		token: null,
		user: null,
		goingTo: null,
	},

	getters: {
		userLoggedIn(state) {
			// TODO check expiration
			// you'll probably have to set timeouts
			return !!state.token
		},
		decodedToken(state) {
			return state.token ? decodeToken(state.token) : null
		},
		userName(state) {
			return state.user ? state.user.name : null
		},
		userId(state, getters) {
			return getters.decodedToken ? getters.decodedToken[0] : null
		},
		userSlug(state) {
			return state.user ? state.user.slug : null
		}
	},

	mutations: {
		login(state, signedUser) {
			const { token, user } = signedUser
			state.token = token
			state.user = user
			privateHttp.defaults.headers.common['Authorization'] = `Bearer ${token}`
			Cookies.set('signedUser', signedUser)
		},

		logout(state) {
			state.token = null
			state.user = null
			delete privateHttp.defaults.headers.common['Authorization']
			Cookies.remove('signedUser')
		},

		setGoingTo(state, location) {
			state.goingTo = location
		},

		unsetGoingTo(state) {
			state.goingTo = null
		},
	},

	actions: {
		grabGoingTo({ commit, state }) {
			const goingTo = state.goingTo
			commit('unsetGoingTo')
			return goingTo
		}
	}
}

export function authPlugin(store) {
	store.watch((state, getters) => getters.userLoggedIn, (isLoggedIn) => {
		if (!isLoggedIn && router.currentRoute.matched.some((route) => route.meta.private)) {
			router.push({ name: 'login' })
		}
	})

	router.beforeEach((to, from, next) => {
		if (to.matched.some(route => route.meta.private)) {
			next({ name: 'login', replace: true })
			store.commit('setGoingTo', to.fullPath)
		}
		else next()
	})
}