<template lang="pug">

#home
	h1 Welcome

	MarkdownEditor(
		v-model="content",
		mode="prosemirror",
	)

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
			content: '**hello**'
		}
	},
}
</script>

<style lang="sass">
</style>
