<template lang="pug">

#project-story
	MarkdownEditor(
		v-model="story",
		ref="markdownEditor",
		:mode="inProsemirror ? 'prosemirror' : 'textarea'",
		:managedImages.sync="managedImages",
		:uploadImage="(image) => console.log(image)",
		:deleteImage="(image) => console.log(image)"
	)

	input(type="checkbox", v-model="inProsemirror")

</template>


<script>
import { sync } from 'vuex-pathify'
import MarkdownEditor from '@/components/MarkdownEditor'

export default {
	components: {
		MarkdownEditor,
	},

	data() {
		return {
			inProsemirror: true,
			managedImages: {},
		}
	},

	computed: {
		...sync('project/', [
			'story',
		]),

		// managedImages: {
		// 	get() {
		// 		return this.project.upload_images.map(imageLink => ({  }))
		// 	},
		// 	set(newManagedImages) {
		// 		// TODO don't bother with this if some are not uploaded

		// 		const managedImages = Object.keys(newManagedImages)
		// 		this.$store.commit('project/updateImages', managedImages)
		// 	},
		// },
	}
}
</script>


<style lang="sass">
</style>
