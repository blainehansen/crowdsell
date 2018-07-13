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
import xxh from 'xxhashjs'

async function sampleHashOfFile(file) {
	return new Promise(function (resolve, reject) {
		const reader = new FileReader()
		const hasher = xxh.h64(0xABCD)

		const oneThird = Math.floor(file.size / 3)
		const twoThird = Math.floor(file.size * 2 / 3)
		const slices = [
			file.slice(0, 256),
			file.slice(oneThird, oneThird + 256),
			file.slice(twoThird, twoThird + 256),
			file.slice(-256),
		]

		reader.onloadend = (event) => {
			if ( event.target.readyState !== FileReader.DONE ) return
			hasher.update(event.target.result)

			if (slices.length > 0) nextSlice()
			else resolve(hasher.digest().toString(36))
		}

		function nextSlice() {
			const slice = slices.pop()
			reader.readAsBinaryString(slice)
		}
		nextSlice()
	})
}

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

			const hash = await sampleHashOfFile(file)
			console.log(hash)
			const { data: urlSlug } = await privateApi.uploadProfilePicture(hash, type, file)
			this.finalUrl = urlSlug
		}
	},
}
</script>

<style lang="sass">
</style>
