<template lang="pug">

.projects
	.project(v-for="project in projects")
		h1 {{ project.title }}
		p {{ project.description }}

</template>

<script>
import { publicApi } from '@/api'

export default {
	name: 'Project',

	props: {
		entitySlug: {
			type: String,
			required: true
		},
		projectSlug: {
			type: String,
			required: true
		},
	},

	asyncData: {
		projects: () => publicApi.getProject()
	}
}
</script>

<style lang="sass">
</style>
