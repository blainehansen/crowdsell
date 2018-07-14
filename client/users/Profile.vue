<template lang="pug">

#profile
	.picture
		input(type="file", accept="image/png, image/jpeg", @change="acceptFile")

		img(v-if="finalUrl", :src="finalUrl | formatSpacesUrl")
		img(v-else-if="previewUrl", :src="previewUrl")

	.location
	.description
	.links

	//- .projects
	//- 	.project(v-for="project in projects")
	//- 		.project-title {{ project.title }}
	//- 		.project-description {{ project.description }}
	//- 		.project-image {{ project.image }}

</template>


<script>
import { privateApi } from '@/api'
import { sampleHashFile } from '@/utils'

export default {
	name: 'profile',

	props: {
		userSlug: {
			type: String,
			default: null,
		}
	},

	data() {
		return {
			finalUrl: null,
			previewUrl: null,
		}
	},
	methods: {
		async acceptFile(event) {
			const file = event.target.files[0]
			this.previewUrl = URL.createObjectURL(file)
			const type = file.type.replace(/^image\//, '')

			const hash = await sampleHashFile(file)
			const { data: urlSlug } = await privateApi.uploadProfilePicture(hash, type, file)
			this.finalUrl = urlSlug
		}
	},

	asyncData: {
		// projects() {
		// 	const userSlug = this.userSlug
		// 	if (!userSlug) return privateApi.getCurrentUserProjects()
		// 	return api.getUserProjects(userSlug)
		// }
	}
}
</script>


<style lang="sass">


</style>
