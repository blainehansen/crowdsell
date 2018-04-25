<template lang="pug">

#profile
	.picture
	.location
	.description
	.links

	.projects
		.project(v-for="project in projects")
			.project-title {{ project.title }}
			.project-description {{ project.description }}
			.project-image {{ project.image }}

</template>


<script>

export default {
	name: 'profile',

	props: {
		userSlug: {
			type: String,
			required: true,
		}
	},

	asyncData: {
		projects() {
			return api.getUserProjects(this.userSlug)
		}
	}
}

</script>


<style lang="sass">


</style>
