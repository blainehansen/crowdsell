<template lang="pug">

.create-project
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
	methods: {
		componentIsActive(componentObject) {
			const active = this.activeComponentObject
			return active && active.id == componentObject.id
		}
	}
}
</script>

<style lang="sass">
</style>
