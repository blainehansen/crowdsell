<template lang="pug">

.projects
	.project(v-for="project in projects")
		h1 {{ project.title }}
		p {{ project.description }}

</template>

<script>
import api from '@/api'

export default {
	name: 'ProjectsIndex',
	asyncData: {
		projects: () => api.getProjects()
	}
}
</script>

<style lang="sass">

</style>
