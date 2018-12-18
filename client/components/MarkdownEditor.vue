<template lang="pug">

div
	.prosemirror(
		ref="prosemirrorEditor",
		v-show="mode === 'prosemirror'",
	)
	template(v-if="mode === 'textarea'")
		textarea.markdown(
			ref="textareaEditor",
			:value="internalMarkdown",
			@input="handleTextareaInput",
		)
		button(@click="insertImageToTextarea") insert image

	input(
		type="file",
		style="display: none",
		ref="filesInput",
		accept="image/png, image/jpeg",
		multiple,
		@change="event => $emit('files', event)",
	)

</template>

<script>

import debounce from 'lodash.debounce'
import isEqual from 'lodash.isequal'

import { sampleHashFile, insertToElement } from '@/utils'
import { schema, managedImagePrefix, markdownSerializerArgs, baseMarkdownParserArgs } from './MarkdownEditorSchema'

import markdownit from 'markdown-it'
import { EditorView } from 'prosemirror-view'
import { exampleSetup } from 'prosemirror-example-setup'
import { EditorState, NodeSelection } from 'prosemirror-state'
import { MarkdownParser, MarkdownSerializer } from 'prosemirror-markdown'
// import { ReplaceStep } from 'prosemirror-transform'

import { MenuItem } from 'prosemirror-menu'
import { buildMenuItems } from 'prosemirror-example-setup'


export default {
	name: 'markdown-editor',

	model: {
		prop: 'markdownValue',
		event: 'markdownChange',
	},

	props: {
		mode: {
			default: 'prosemirror',
			type: String,
			required: false,
			validator: value => ['prosemirror', 'textarea'].includes(value),
		},
		markdownValue: String,

		managedImages: {
			type: Object,
			validator(object) {
				for (const [key, value] of Object.entries(object)) {
					if (typeof key !== 'string') return false
					if (!(value.file instanceof File)) return false
					if (typeof value.url !== 'string') return false
					if (typeof value.key !== 'string') return false
					if (typeof value.uploaded !== 'boolean') return false
				}

				return true
			},
		},

		uploadImage: Function,
		deleteImage: Function,
	},

	data() {
		return {
			internalMarkdown: '',

			markdownParserInstance: null,
			markdownSerializerInstance: null,
		}
	},

	mounted() {
		const menu = buildMenuItems(schema)
		const existingInsertImage = menu.insertImage.spec

		menu.insertMenu.content[0] = menu.insertImage = new MenuItem({
			...existingInsertImage,
			label: "Image",

			run: async (state, _, view) => {
				const { from: $from, to } = state.selection

				// this is allowing editing of existing images
				// if we're already over an image, we grab that node's attrs and use them to populate the thing
				const attrs = state.selection instanceof NodeSelection && state.selection.node.type == schema.nodes.image
					? state.selection.node.attrs
					// this grabs the selected text, if there is any
					// for the situation where they aren't over an image node
					: { alt: state.doc.textBetween($from, to, " ") }

				// deploy the picker
				const [file,] = await this.getFiles()
				const managedImage = await this.makeAndManageImage(file)
				attrs.key = managedImage.key
				attrs.src = managedImage.url
				attrs.title = managedImage.file.name

				// here's the real code that actually applies the thing
				view.dispatch(view.state.tr.replaceSelectionWith(schema.nodes.image.createAndFill(attrs)))
				view.focus()
			}
		})
		this.menu = menu

		this.markdownSerializerInstance = new MarkdownSerializer(...markdownSerializerArgs)

		const markdownParserArgs = {
			...baseMarkdownParserArgs,
			image: {
				...baseMarkdownParserArgs.image,
				getAttrs: tok => {
					const foundSrc = tok.attrGet("src")
					const foundManaged = this.managedImages[foundSrc]

					return {
						key: foundManaged ? foundManaged.key : null,
						src: foundManaged ? foundManaged.url : foundSrc,
						alt: tok.content || null,
						title: tok.attrGet("title") || null,
					}
				}
			},
		}
		this.markdownParserInstance = new MarkdownParser(schema, markdownit("commonmark", { html: false }), markdownParserArgs)

		const internalMarkdown = this.internalMarkdown = this.markdownValue

		this.view = new EditorView(this.$refs.prosemirrorEditor, {
			state: this.createState(internalMarkdown),

			dispatchTransaction: transaction => {
				this.view.updateState(this.view.state.apply(transaction))

				if (!transaction.docChanged || transaction.steps.length === 0) return
				// TODO maybe consider making this function debounced?
				this.applyMarkdown(this.markdownSerializerInstance.serialize(this.view.state.doc))
				this.syncPictures()
			},
		})

		if (this.mode === 'prosemirror') this.view.focus()
		else this.$refs.textareaEditor.focus()
	},

	methods: {
		syncPictures: debounce(function() {
			// ensure the view is up to date
			if (this.mode !== 'prosemirror') this.rebuildProsemirror()

			const foundKeys = new Set()
			this.view.state.doc.descendants(node => {
				// TODO some sort of filter to say if it can hold an image
				if (node.type.name !== 'image') return node.type.isBlock

				if (node.attrs.key.startsWith(managedImagePrefix))
					foundKeys.add(node.attrs.key)

				return false
			})

			const managedImages = this.managedImages
			const newManagedImages = {}

			for (const [key, value] of Object.entries(managedImages)) {
				if (foundKeys.has(key))
					newManagedImages[key] = value

				else if (!value.uploaded)
					URL.revokeObjectURL(value.url)
			}

			this.$emit('update:managedImages', newManagedImages)
		}, 1000),


		getFiles() {
			return new Promise(resolve => {
				this.$once('files', event => {
					const eventFiles = event.target.files
					const files = []
					for (let index = eventFiles.length - 1; index >= 0; index--)
						files.push(eventFiles[index])

					this.$refs.filesInput.value = null
					resolve(files)
				})

				this.$refs.filesInput.click()
			})
		},

		async makeAndManageImage(file) {
			const hash = await sampleHashFile(file)
			const key = managedImagePrefix + hash
			const url = URL.createObjectURL(file)

			const existing = this.managedImages[key]
			const newManagedImages = this.managedImages
			if (existing) {
				URL.revokeObjectURL(existing.url)
			}
			const newImage = { file, url, key, hash, uploaded: false }
			newManagedImages[key] = newImage
			this.$emit('update:managedImages', newManagedImages)

			return newImage
		},

		async insertImageToTextarea() {
			const [file,] = await this.getFiles()

			const managedImage = await this.makeAndManageImage(file)

			// push the string representation to the editor
			const imageString = `![alt text](${managedImage.key} "${managedImage.file.name}")`
			insertToElement(this.$refs.textareaEditor, imageString)
		},

		handleTextareaInput(event) {
			this.applyMarkdown(event.target.value)
			this.syncPictures()
		},

		applyMarkdown(newMarkdown) {
			this.internalMarkdown = newMarkdown
			this.$emit('markdownChange', newMarkdown)
		},

		createState(newMarkdown) {
			return EditorState.create({
				doc: this.markdownParserInstance.parse(newMarkdown),
				plugins: exampleSetup({ schema: schema, menuContent: this.menu.fullMenu })
			})
		},

		rebuildProsemirror() {
			const state = this.createState(this.internalMarkdown)
			this.view.updateState(state)
		}
	},

	watch: {
		markdownValue(newMarkdownValue) {
			if (newMarkdownValue === this.internalMarkdown) return

			this.internalMarkdown = newMarkdownValue

			if (this.mode === 'prosemirror') this.rebuildProsemirror()
		},

		managedImages(newManagedImages) {
			if (isEqual(newManagedImages, this.managedImages)) return
			if (this.mode === 'prosemirror') this.rebuildProsemirror()
		},

		mode(newMode, oldMode) {
			if (newMode !== oldMode && newMode === 'prosemirror') {
				this.rebuildProsemirror()
			}
		}
	}
}

</script>

<style lang="sass">

/* from Prosemirror */
.ProseMirror
	position: relative

.ProseMirror
	white-space: pre-wrap

.ProseMirror ul, .ProseMirror ol
	padding-left: 30px
	cursor: default

.ProseMirror blockquote
	padding-left: 1em
	border-left: 3px solid #eee
	margin-left: 0
	margin-right: 0

.ProseMirror pre
	white-space: pre-wrap

.ProseMirror li
	position: relative
	/* Don't do weird stuff with marker clicks */
	pointer-events: none

.ProseMirror li > *
	pointer-events: auto

.ProseMirror-nodeselection *::selection, .ProseMirror-widget *::selection
	background: transparent

.ProseMirror-nodeselection *::-moz-selection, .ProseMirror-widget *::-moz-selection
	background: transparent

.ProseMirror-selectednode
	outline: 2px solid #8cf

/* Make sure li selections wrap around markers */
li.ProseMirror-selectednode
	outline: none

li.ProseMirror-selectednode:after
	content: ""
	position: absolute
	left: -32px
	right: -2px
	top: -2px
	bottom: -2px
	border: 2px solid #8cf
	pointer-events: none

/* from markdown example */
.ProseMirror
	height: 120px
	overflow-y: auto
	box-sizing: border-box
	-moz-box-sizing: border-box

textarea
	width: 100%
	height: 123px
	border: 1px solid silver
	box-sizing: border-box
	-moz-box-sizing: border-box
	padding: 3px 10px
	border-radius: 3px
	border: 1px solid #38a

.ProseMirror-menubar-wrapper, #markdown textarea
	display: block
	margin-bottom: 4px

/* from menu bar */
.ProseMirror-textblock-dropdown
	min-width: 3em

.ProseMirror-menu
	margin: 0 -4px
	line-height: 1

.ProseMirror-tooltip .ProseMirror-menu
	width: -webkit-fit-content
	width: fit-content
	white-space: pre

.ProseMirror-menuitem
	margin-right: 3px
	display: inline-block

.ProseMirror-menuseparator
	border-right: 1px solid #ddd
	margin-right: 3px

.ProseMirror-menu-dropdown, .ProseMirror-menu-dropdown-menu
	font-size: 90%
	white-space: nowrap

.ProseMirror-menu-dropdown
	vertical-align: 1px
	cursor: pointer

.ProseMirror-menu-dropdown-wrap
	padding: 1px 14px 1px 4px
	display: inline-block
	position: relative

.ProseMirror-menu-dropdown:after
	content: ""
	border-left: 4px solid transparent
	border-right: 4px solid transparent
	border-top: 4px solid currentColor
	opacity: .6
	position: absolute
	right: 2px
	top: calc(50% - 2px)

.ProseMirror-menu-dropdown-menu, .ProseMirror-menu-submenu
	position: absolute
	background: white
	color: #666
	border: 1px solid #aaa
	padding: 2px

.ProseMirror-menu-dropdown-menu
	z-index: 15
	min-width: 6em

.ProseMirror-menu-dropdown-item
	cursor: pointer
	padding: 2px 8px 2px 4px

.ProseMirror-menu-dropdown-item:hover
	background: #f2f2f2

.ProseMirror-menu-submenu-wrap
	position: relative
	margin-right: -4px

.ProseMirror-menu-submenu-label:after
	content: ""
	border-top: 4px solid transparent
	border-bottom: 4px solid transparent
	border-left: 4px solid currentColor
	opacity: .6
	position: absolute
	right: 4px
	top: calc(50% - 4px)

.ProseMirror-menu-submenu
	display: none
	min-width: 4em
	left: 100%
	top: -3px

.ProseMirror-menu-active
	background: #eee
	border-radius: 4px

.ProseMirror-menu-active
	background: #eee
	border-radius: 4px

.ProseMirror-menu-disabled
	opacity: .3

.ProseMirror-menu-submenu-wrap:hover .ProseMirror-menu-submenu, .ProseMirror-menu-submenu-wrap-active .ProseMirror-menu-submenu
	display: block

.ProseMirror-menubar
	border-top-left-radius: inherit
	border-top-right-radius: inherit
	position: relative
	min-height: 1em
	color: #666
	padding: 1px 6px
	top: 0
	left: 0
	right: 0
	border-bottom: 1px solid silver
	z-index: 10
	-moz-box-sizing: border-box
	box-sizing: border-box
	overflow: visible

.ProseMirror-icon
	display: inline-block
	line-height: .8
	/* Compensate for padding */
	vertical-align: -2px
	padding: 2px 8px
	cursor: pointer

.ProseMirror-icon svg
	fill: currentColor
	height: 1em

.ProseMirror-icon span
	vertical-align: text-top

.ProseMirror-prompt
	background: white
	padding: 5px 10px 5px 15px
	border: 1px solid silver
	position: fixed
	border-radius: 3px
	z-index: 11
	box-shadow: -.5px 2px 5px rgba(0, 0, 0, .2)

.ProseMirror-prompt h5
	margin: 0
	font-weight: normal
	font-size: 100%
	color: #444

.ProseMirror-prompt input[type="text"],
.ProseMirror-prompt textarea
	background: #eee
	border: none
	outline: none

.ProseMirror-prompt input[type="text"]
	padding: 0 4px

.ProseMirror-prompt-close
	position: absolute
	left: 2px
	top: 1px
	color: #666
	border: none
	background: transparent
	padding: 0

.ProseMirror-prompt-close:after
	content: "✕"
	font-size: 12px

.ProseMirror-invalid
	background: #ffc
	border: 1px solid #cc7
	border-radius: 4px
	padding: 5px 10px
	position: absolute
	min-width: 10em

.ProseMirror-prompt-buttons
	margin-top: 5px
	display: none

textarea.vue-prosemirror,
.vue-prosemirror div.ProseMirror-content
	border: none
	overflow: auto
	outline: none
	-webkit-box-shadow: none
	-moz-box-shadow: none
	box-shadow: none

textarea.vue-prosemirror
	margin: 0
	padding: 0


</style>
