<template lang="pug">

.projects
	.project(v-for="project in projects")
		h1 {{ project.title }}
		p {{ project.description }}

</template>

<script>
import { publicApi } from '@/api'

export default {
	name: 'ProjectsIndex',
	asyncData: {
		projects: () => publicApi.getProjects()
	}
}
</script>

<style lang="sass">
.thing
	color: red

</style>
