<template lang="pug">

.create-project
	p
		button(v-if="project$anyTouched", @click="saveProject") save project
		button(v-else, disabled) project saved

	ul.nav.justify-content-center
		li.nav-item(v-for="componentObject in componentManifest")
			.nav-link(
				@click="activeComponentObject = componentObject",
				:class="{ active: componentIsActive(componentObject) }"
			) {{ componentObject.pageName }}

	template(v-if="activeComponentObject")
		h1 {{ activeComponentObject.componentTitle }}
		p {{ activeComponentObject.componentDescription }}
		component(v-if="activeComponentObject", :is="activeComponentObject.component")

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

import Overall from './create/Overall'

const componentManifest = [
	{ id: 1, component: Overall, pageName: "Overall", componentTitle: "Get started", componentDescription: "Make the overall decisions" },
]

export default {
	name: 'CreateProject',
	data() {
		return {
			componentManifest,
			activeComponentObject: null,
		}
	},
	computed: {
		project$anyTouched: get('project/$anyTouched'),
	},

	methods: {
		componentIsActive(componentObject) {
			const active = this.activeComponentObject
			return active && active.id == componentObject.id
		},

		saveProject: call('project/saveProject'),
	}
}
</script>

<style lang="sass">
</style>
