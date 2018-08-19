import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import Login from '@/Login'

// import ProjectsIndex from '@/projects/ProjectsIndex'
import CreateProject from '@/projects/CreateProject'
import Project from '@/projects/Project'

import CreateProjectManifest from '@/projects/create'

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

		// {
		// 	path: '/projects/create', component: CreateProject,
		// 	meta: { private: true },
		// },

		{
			path: '/projects/create', component: CreateProject,
			props: { projectId: null },
			meta: { private: true },
			children: CreateProjectManifest,
		},

		// {
		// 	path: '/projects/create/:projectId', component: EditProject,
		// 	props: true,
		// 	meta: { private: true },
		// 	children: CreateProjectManifest,
		// },
		// { path: '/projects/:userSlug/:projectSlug', name: 'project', component: Project, props: true },

		{ path: '/you', name: 'userProfile', component: UserProfile, meta: { private: true } },
		{ path: '/profile/:userSlug', name: 'profile', component: Profile, props: { previewing: false } },

		{ path: '/about', name: 'about', component: About },

		// a blog
		// an engineering blog?
	]
})

export default router
