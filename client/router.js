import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import Login from '@/Login'

import ProjectsIndex from '@/projects/ProjectsIndex'
import NewProject from '@/projects/NewProject'
import Project from '@/projects/Project'

import UserProfile from '@/users/UserProfile'
import Profile from '@/users/Profile'

import About from '@/pages/About'

Vue.use(Router)

const router = new Router({
	mode: 'history',
	routes: [
		{ path: '/', name: 'home', component: Home },
		{ path: '/login', name: 'login', component: Login },

		{ path: '/projects', name: 'projects', component: ProjectsIndex },

		{ path: '/projects/new', name: 'projectsNew', component: NewProject },
		{ path: '/projects/:userSlug/:projectSlug', name: 'project', component: Project, props: true },

		{ path: '/profile', name: 'currentUserProfile', component: Profile, props: { userSlug: null, projectSlug: null } },

		{ path: '/profile/:userSlug', name: 'userProfile', component: Profile, props: true },

		{ path: '/about', name: 'about', component: About },

		// a blog
		// an engineering blog?
	]
})

export default router

import Cookies from 'js-cookie'
import api, { privateHttp } from '@/api'

async function decodeTokenToUser(token) {
	const segments = token.split('.')
	if (segments.length != 2) console.error('invalid token', token)
	const [id, , ,] = JSON.parse(atob(segments[0]))

	const { data: user } = await api.getUserInfo(id)
	return user
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
		userName(state) {
			return state.user ? state.user.name : null
		}
	},

	mutations: {
		async login(state, token) {
			state.token = token
			privateHttp.defaults.headers.common['Authorization'] = `Bearer ${token}`
			Cookies.set('authToken', token)

			state.user = await decodeTokenToUser(token)
		},

		logout(state) {
			state.token = null
			state.user = null
			delete privateHttp.defaults.headers.common['Authorization']
			Cookies.remove('authToken')
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
	// store.watch((state, getters) => getters.userLoggedIn, (isLoggedIn) => {
	// 	if (!isLoggedIn) {
	// 		router.push({ name: 'login' })
	// 	}
	// })

	router.beforeEach((to, from, next) => {
		if (to.matched.some(record => record.meta.private)) {
			next({ name: 'login', replace: true })
			store.commit('setGoingTo', to.fullPath)
		}
		else next()
	})
}
