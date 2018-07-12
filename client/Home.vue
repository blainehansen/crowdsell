<template lang="pug">

#home
	h1 Welcome

	input(type="file", accept="image/png, image/jpeg", @change="acceptFile")

	img(v-if="finalUrl", :src="createFinalUrl(finalUrl)")
	img(v-else-if="previewUrl", :src="previewUrl")

	//- MdEditor(v-model="myText")

	//- p {{ myText }}

</template>

<script>
// import MdEditor from '@/components/MdEditor'

import { privateApi } from '@/api'
import { delay } from 'bluebird'
import axios from 'axios'

export default {
	name: 'home',
	// components: {
	// 	MdEditor
	// },
	// data() {
	// 	return { myText: '## hello' }
	// }
	data() {
		return {
			finalUrl: null,
			previewUrl: null,
		}
	},
	methods: {
		createFinalUrl(finalUrl) {
			return `https://blaine-final-spaces-test.nyc3.cdn.digitaloceanspaces.com/${finalUrl}`
		},
		async acceptFile(event) {
			const file = event.target.files[0]
			this.previewUrl = URL.createObjectURL(file)
			const type = file.type.replace(/^image\//, '')
			// TODO this isn't correct
			const hash = btoa(file.name).replace(/\=/g, '')
			console.log(hash)
			const { data: urlSlug } = await privateApi.uploadProfilePicture(hash, type, file)

			console.log(urlSlug)
			this.finalUrl = urlSlug
			// const uploadedFile = event.target.files
			// for (let i = 0; i < uploadedFiles.length; i++) {
			// 	files.push(uploadedFiles[i])
			// }
			// this.currentFileName = files[0].name
			// this.$emit('files', files)

			// this.$refs.files.value = null
		}
	},
}
</script>

<style lang="sass">
</style>
