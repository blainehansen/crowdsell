<template lang="pug">

.edit-project
	p
		button(v-if="project$anyTouched", @click="saveProject") save project
		button(v-else, disabled) project saved

	ul.nav.justify-content-center
		li.nav-item(v-for="componentObject in componentManifest")
			router-link.nav-link(:to="{ name: componentObject.name }") {{ componentObject.pageName }}

		//- router-link.nav-item(tag="li", :to="{ name: componentObject.name }", v-for="componentObject in componentManifest")
		//- 	a.nav-link {{ componentObject.pageName }}

	router-view

</template>

<script>
import { call, get } from 'vuex-pathify'

import componentManifest from './edit'

export default {
	name: 'EditProject',

	props: {
		projectId: {
			type: String,
			required: true,
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
