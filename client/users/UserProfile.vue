<template lang="pug">

#user
	p hello user

	.picture
		input(type="file", accept="image/png, image/jpeg", @change="acceptFile")

		img(v-if="finalUrl", :src="finalUrl | formatSpacesUrl")
		img(v-else-if="previewUrl", :src="previewUrl")

	.name
		input(v-model="name")
	.location
		input(v-model="locationSearch")
	.description
		input(v-model="bio")
	.links
		.link(v-for="link in user.links")
			input(v-model="link")
			//- a delete button
		//- an add link button
		.link(v-if="addingLink")
			input(v-model="newLink")

	.custom-slug
		//- changing this should directly call a mutation
		input(v-model="customSlug")

	//- .projects
	//- 	.project(v-for="project in projects")
	//- 		.project-title {{ project.title }}
	//- 		.project-description {{ project.description }}
	//- 		.project-image {{ project.image }}

</template>


<script>
import { privateApi } from '@/api'
import { sampleHashFile } from '@/utils'
import { call, get, sync } from '@/packages/vuex-pathify'

export default {
	name: 'userProfile',
	data() {
		return {
			finalUrl: null,
			previewUrl: null,
		}
	},

	beforeCreate() {
		this.$store.dispatch('user/fetchInitial')
	},

	computed: {
		slug: get('auth/userSlug'),
		...sync('user@', ['name', 'bio', 'links', 'location']),
	},

	methods: {
		async changeSlug(newSlug) {
			const { data: signedUser }= await privateApi.changeSlug(this.slug, newSlug)
			this.$store.commit('auth/login', signedUser)
		},

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
