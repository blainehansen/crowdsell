import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import Login from '@/Login'

// import ProjectsIndex from '@/projects/ProjectsIndex'
import CreateProject from '@/projects/CreateProject'
import EditProject from '@/projects/EditProject'
// import Project from '@/projects/Project'

import editProjectSteps from '@/projects/edit'

import UserProfile from '@/users/UserProfile'
import Profile from '@/users/Profile'

import About from '@/pages/About'

Vue.use(Router)

const router = new Router({
	mode: 'history',

	linkActiveClass: 'active',

	routes: [
		{
			path: '/',
			name: 'home',
			component: Home,
		},
		{
			path: '/login',
			name: 'login',
			component: Login,
		},

		// { path: '/projects', name: 'projects', component: ProjectsIndex },

		{
			path: '/projects/create',
			name: 'projectCreate',
			component: CreateProject,
			meta: { private: true },
		},

		{
			path: '/projects/create/:projectId',
			component: EditProject,
			props: true,
			meta: { private: true },
			children: editProjectSteps,
		},

		// { path: '/projects/:userSlug/:projectSlug', name: 'project', component: Project, props: true },

		{
			path: '/you',
			name: 'userProfile',
			component: UserProfile,
			meta: { private: true },
		},

		{
			path: '/you/preview',
			name: 'userProfilePreview',
			component: Profile,
			props: { previewing: true },
			meta: { private: true },
		},
		{
			path: '/profile/:userSlug',
			name: 'profile',
			component: Profile,
			props: { previewing: false }
		},

		{
			path: '/about',
			name: 'about',
			component: About,
		},

		// a blog
		// an engineering blog?
	]
})

export default router


// TODO https://alligator.io/vuejs/vue-router-modify-head/
