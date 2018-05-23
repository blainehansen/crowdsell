<template lang="pug">

#home

	h1 Welcome to Crowd Sell

	.projects
		.project(v-for="project in projects")
			h1 {{ project.title }}
			p {{ project.description }}

</template>

<script>
import api from '@/api'

export default {
	name: 'home',
	asyncData: {
		projects: () => api.getProjects()
	}
}
</script>

<style lang="sass">

</style>
