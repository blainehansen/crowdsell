<template lang="pug">

#home
	h1 Welcome

	MarkdownEditor(
		v-model="content",
		ref="markdownEditor",
		:mode="inProsemirror ? 'prosemirror' : 'textarea'",
		:managedImages.sync="managedImages",
		:uploadImage="(image) => { console.log(image); return image.url }",
		:deleteImage="(image) => console.log(image)"
	)

	input(type="checkbox", v-model="inProsemirror")


</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'

export default {
	name: 'home',

	components: {
		MarkdownEditor,
	},

	data() {
		return {
			content: "**hello**",
			inProsemirror: true,
			internalManagedImages: {},
		}
	},

	computed: {
		managedImages: {
			get() {
				return this.internalManagedImages
			},
			set(newManagedImages) {
				// TODO don't bother with this if some are not uploaded
				this.internalManagedImages = newManagedImages
			},
		}
	}
}

</script>

<style lang="sass">
</style>
