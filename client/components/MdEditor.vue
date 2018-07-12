<template lang="pug">

#editor

</template>

<script>
import { default as Editor } from 'tui-editor'

require('codemirror/lib/codemirror.css')
require('tui-editor/dist/tui-editor.css')
require('tui-editor/dist/tui-editor-contents.css')
require('highlight.js/styles/github.css')

export default {
	name: 'MdEditor',
	props: {
		value: {
			type: String,
			required: true,
		}
	},
	mounted() {
		const editor = new Editor({
			el: this.$el,
			initialValue: this.value,
			height: 'auto',
			usageStatistics: false,
			previewStyle: 'vertical',
			initialEditType: 'wysiwyg',
			events: {
				change: () => {
					this.$emit('input', editor.getMarkdown())
				}
			},
		})

		this.$editor = editor
	},
}
</script>

<style lang="sass">
</style>
