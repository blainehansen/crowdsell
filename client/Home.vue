<template lang="pug">

#home
	h1 Welcome

	MarkdownEditor(
		initialMarkdown="**hello**",
		mode="all",
		:textareaConfig="{name: 'my-textarea'}",
		@contentChangeMarkdown="markdownChangedHandler",
	)

</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'

export default {
	name: 'home',

	components: {
		MarkdownEditor,
	},
}
</script>

<style lang="sass">
</style>
