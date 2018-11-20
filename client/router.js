import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import Login from '@/Login'

import CreateProject from '@/projects/CreateProject'
import EditProject from '@/projects/EditProject'
import editProjectSteps from '@/projects/edit'
import Project from '@/projects/Project'

import ProjectsExplore from '@/projects/ProjectsExplore'

import ProjectPledge from '@/projects/ProjectPledge'
import ProjectConfirm from '@/projects/ProjectConfirm'

import UserProfile from '@/users/UserProfile'
import Profile from '@/users/Profile'

import About from '@/pages/About'

Vue.use(Router)

const router = new Router({
	mode: 'history',

	linkActiveClass: 'active',
	linkExactActiveClass: 'active-exact',

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

		{
			path: '/projects',
			name: 'projectsExplore',
			component: ProjectsExplore,
			props: route => ({
				query: route.query.q,
				categories: route.query.c,
				tags: route.query.t,
			}),
		},

		// individual project
		{
			path: '/projects/:userSlug/:projectSlug',
			name: 'project',
			component: Project,
			props: true,
		},

		// payment flow
		{
			path: '/projects/:userSlug/:projectSlug/pledge',
			name: 'projectPledge',
			component: ProjectPledge,
			props: true,
		},

		// feedback flow
		{
			path: '/projects/:userSlug/:projectSlug/confirm',
			name: 'projectConfirm',
			component: ProjectConfirm,
			props: true,
		},

		// start project
		{
			path: '/projects/create',
			name: 'projectCreate',
			component: CreateProject,
			meta: { private: true },
		},

		// edit project
		{
			path: '/projects/create/:projectId',
			component: EditProject,
			props: true,
			meta: { private: true },
			children: editProjectSteps,
		},


		// your profile
		{
			path: '/you',
			name: 'userProfile',
			component: UserProfile,
			meta: { private: true },
		},

		// your profile preview
		{
			path: '/you/preview',
			name: 'userProfilePreview',
			component: Profile,
			props: { previewing: true },
			meta: { private: true },
		},

		// someone else's profile
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
