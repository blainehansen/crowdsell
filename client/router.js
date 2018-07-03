import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import Login from '@/Login'

// import ProjectsIndex from '@/projects/ProjectsIndex'
import CreateProject from '@/projects/CreateProject'
import Project from '@/projects/Project'

import UserProfile from '@/users/UserProfile'
import Profile from '@/users/Profile'

import About from '@/pages/About'

Vue.use(Router)

const router = new Router({
	mode: 'history',

	linkActiveClass: 'active',

	routes: [
		{ path: '/', name: 'home', component: Home },
		{ path: '/login', name: 'login', component: Login },

		// { path: '/projects', name: 'projects', component: ProjectsIndex },

		{ path: '/projects/create', name: 'projectCreate', component: CreateProject, meta: { private: true } },
		// { path: '/projects/:userSlug/:projectSlug', name: 'project', component: Project, props: true },

		// { path: '/profile', name: 'currentUserProfile', component: Profile, props: { userSlug: null, projectSlug: null } },

		// { path: '/profile/:userSlug', name: 'userProfile', component: Profile, props: true },

		{ path: '/about', name: 'about', component: About },

		// a blog
		// an engineering blog?
	]
})

export default router
