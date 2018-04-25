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

export default new Router({
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


import Cookies from 'js-cookie'

export const auth = {
	state: {
		locationInApp: null,
		user: {
			name: null,
			expiresIn: 0,
			accessToken: null,
			tokenCreationTimestamp: 0,
		}
	},
	getters: {
		loggedIn(state) {
			// we need to make sure claims are signed with a valid public key,
			// and we need to use the exp claim because it's correctly signed
			// https://tools.ietf.org/html/rfc7519#section-4.1.4
			const timestamp = moment().unix()
			return !!state.user.accessToken && (timestamp < state.user.tokenCreationTimestamp + state.user.expiresIn)
		},
	},
	mutations: {
		setLocationInApp(state, name) {
			state.locationInApp = name
		},

		loginUser(state, {accessToken, userObj, expiresIn, tokenCreationTimestamp}) {
			const userObject = {
				name: userObj.name,
				accessToken: accessToken,
				expiresIn: expiresIn || 86400,
				tokenCreationTimestamp: tokenCreationTimestamp
			}

			state.user = cloneDeep(userObject)
			axios.defaults.headers.common['Authorization'] = `Bearer ${accessToken}`

			Cookies.set('authenticatedUser', userObject)

			// TODO settimeout
		},

		logoutUser(state) {
			state.user = { name: null, accessToken: null, expiresIn: 0, tokenCreationTimestamp: 0 }
			axios.defaults.headers.common['Authorization'] = null

			Cookies.remove('authenticatedUser')
		}
	}
}


export function authPlugin(store) {
	store.watch((state, getters) => getters.loggedIn, (isLoggedIn) => {
		if (!isLoggedIn) {
			router.push({ name: 'login' })
		}
	})

	router.beforeEach((to, from, next) => {
		if (store.getters.loggedIn) {
			next()
		}
		else if (!to.meta.allowAnonymous) {
			next({ name: 'login', replace: true })
			store.commit('setLocationInApp', to.name)
		}
		else {
			next()
		}
	})
}
