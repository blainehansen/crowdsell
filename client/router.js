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

		{
			path: '/projects/create', name: 'projectCreate', component: CreateProject,
			props: { projectId: null },
			meta: { private: true },
			children: CreateProjectManifest,
		},

		{
			path: '/projects/create/:projectId', name: 'projectEdit', component: CreateProject,
			props: true,
			meta: { private: true },
			children: CreateProjectManifest,
		}
		// { path: '/projects/:userSlug/:projectSlug', name: 'project', component: Project, props: true },

		// { path: '/profile', name: 'currentUserProfile', component: Profile, props: { userSlug: null } },
		// { path: '/profile/:userSlug', name: 'userProfile', component: Profile, props: true },

		{ path: '/about', name: 'about', component: About },

		// a blog
		// an engineering blog?
	]
})

export default router
