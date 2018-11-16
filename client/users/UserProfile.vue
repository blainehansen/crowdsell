<template lang="pug">

#user
	p hello user

	.picture
		input(type="file", accept="image/png, image/jpeg", @change="acceptFile")

		img(v-if="finalVersion", :src="finalVersion | formatProfileImageUrl")
		img(v-else-if="previewUrl", :src="previewUrl")

	p
		button(v-if="user$anyTouched", @click="saveUser") save changes
		button(v-else, disabled) saved

	.name
		input(v-model="name", placeholder="name")
	.location
		input(v-model="location", placeholder="location")
	.description
		input(v-model="bio", placeholder="bio")
	.links
		input(v-model="links", placeholder="links")
		//- .link(v-for="link in links")
		//- 	input(v-model="link")
		//- 	//- a delete button
		//- //- an add link button
		//- .link(v-if="addingLink")
		//- 	input(v-model="newLink")

	//- .custom-slug
	//- 	input(v-model="customSlug")

	//- .projects
	//- 	.project(v-for="project in projects")
	//- 		.project-title {{ project.title }}
	//- 		.project-description {{ project.description }}
	//- 		.project-image {{ project.image }}

</template>


<script>
import { secureApi, imagesApi } from '@/api'
// import { sampleHashFile } from '@/utils'
import { call, get, sync } from 'vuex-pathify'

export default {
	name: 'userProfile',
	data() {
		return {
			finalVersion: null,
			previewUrl: null,
		}
	},

	beforeCreate() {
		this.$store.dispatch('user/fetchInitial')
	},

	computed: {
		slug: get('auth/userSlug'),
		...sync('user/', ['name', 'bio', 'links', 'location']),
		user$anyTouched: get('user/$anyTouched'),
	},

	methods: {
		async changeSlug(newSlug) {
			const { data: signedUser }= await secureApi.changeSlug(newSlug)
			this.$store.commit('auth/login', signedUser)
		},

		async acceptFile(event) {
			const file = event.target.files[0]
			this.previewUrl = URL.createObjectURL(file)

			const version = await imagesApi.uploadProfileImage(file)

			this.finalVersion = version
		},

		saveUser: call('user/saveUser'),
	},
}

</script>


<style lang="sass">
</style>
