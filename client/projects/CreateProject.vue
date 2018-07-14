<template lang="pug">

.create-project
	p
		button(v-if="project$anyTouched", @click="saveProject") save project
		button(v-else, disabled) project saved

	ul.nav.justify-content-center
		li.nav-item(v-for="componentObject in componentManifest")
			router-link.nav-link(:to="{ name: componentObject.name }") {{ componentObject.pageName }}

		//- router-link.nav-item(tag="li", :to="{ name: componentObject.name }", v-for="componentObject in componentManifest")
		//- 	a.nav-link {{ componentObject.pageName }}

	//- portal-target(name="createProjectDescription")
	router-view

	h2(v-else) Choose a step


//- title
//- blurb
//- image
//- video?
//- categories and tags
//- funding goal
//- funding timeline and structure
//- add people to project
//- description
//- the files themselves

//- the demo materials
//- the release guarantees
//- the license

//- financial information

</template>

<script>
import api from '@/api'
import { call, get } from '@/packages/vuex-pathify'

import componentManifest from './create'

export default {
	name: 'CreateProject',

	props: {
		projectId: {
			type: Number,
			default: null,
		}
	},

	data() {
		return {
			componentManifest,
		}
	},
	computed: {
		project$anyTouched: get('project/$anyTouched'),
	},

	methods: {
		saveProject: call('project/saveProject'),
	}
}
</script>

<style lang="sass">
</style>
