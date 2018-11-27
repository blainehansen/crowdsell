<template lang="pug">

div
	div
		input(type="file", accept="image/png, image/jpeg", multiple, @change="acceptFiles")

	//- so there's two views
	//- one that's a bare bones textarea that just uses v-model
	//- the other is a fully enabled prosemirror instance

	//- when we're in markdown mode, we just handle v-model events like normal
	//- when we're in prosemirror mode, we hook into dispatchTransaction

	.prosemirror(
		ref="prosemirrorEditor",
		v-show="mode === 'prosemirror'",
	)
	textarea.markdown(
		v-show="mode === 'textarea'",
		:value="internalMarkdown",
		@input="applyMarkdown",
	)
	//- button(@click="insertTextAreaImage")

</template>

<script>

import { sampleHashFile } from '@/utils'
import markdownit from "markdown-it"
import { EditorView } from 'prosemirror-view'
import { EditorState } from 'prosemirror-state'
import { exampleSetup } from 'prosemirror-example-setup'
// MarkdownSerializer
import { schema, MarkdownParser, defaultMarkdownSerializer } from 'prosemirror-markdown'

const managedImages = {}
const managedImagePrefix = 'managed-image:'
let foundKeys = new Set()

async function addManagedImage(file) {
	const hash = await sampleHashFile(file)
	const key = managedImagePrefix + hash

	const existing = managedImages[key]
	if (existing) {
		URL.revokeObjectURL(existing.url)
	}
	const managedImage = {
		key,
		name: file.name,
		url: URL.createObjectURL(file),
	}
	managedImages[key] = managedImage
	return managedImage
}


const defaultMarkdownParser = new MarkdownParser(schema, markdownit("commonmark", { html: false }), {
	blockquote: { block: "blockquote" },
	paragraph: { block: "paragraph" },
	list_item: { block: "list_item" },
	bullet_list: { block: "bullet_list" },
	ordered_list: { block: "ordered_list", getAttrs: tok => ({ order: +tok.attrGet("order") || 1 }) },
	heading: { block: "heading", getAttrs: tok => ({ level: +tok.tag.slice(1) }) },
	code_block: { block: "code_block" },
	fence: { block: "code_block", getAttrs: tok => ({ params: tok.info || "" }) },
	hr: { node: "horizontal_rule" },
	image: {
		node: "image",
		getAttrs(tok) {
			const alt = tok.content || ""
			if (alt.startsWith(managedImagePrefix)) {
				foundKeys.add(alt)
			}

			return {
				alt: alt || null,
				src: tok.attrGet("src"),
				title: tok.attrGet("title") || null,
			}
		}
	},

	hardbreak: { node: "hard_break" },

	em: { mark: "em" },
	strong: { mark: "strong" },
	link: { mark: "link", getAttrs: tok => ({
		href: tok.attrGet("href"),
		title: tok.attrGet("title") || null,
	})},
	code_inline: { mark: "code" }
})

// we always maintain a list of images in the document managed by the site
// when the upload manager adds them, it creates a hash and prefixes it with 'managed-upload' or something
// it's an object mapping these prefixed hashes to actual data or url's or whatever

// the image uploader also has a "from url option" if someone wants to do that

// in raw markdown mode, there's a button that can add images. it takes the file, adds it to the managed list, and inserts a markdown snippet at the current cursor position. when this markdown string is viewed as html, we just append a "footnote" area to the string containing footnotes for all the images. this gives us the capability to have file url's or whatever act as

// since we prefix managed ones, we can detect when we parse (serlialize?) if any have been messed with or broken in some way

// from a performance perspective, the best thing is to not touch the api until they persist changes
// then we do a diff of the hashes already added to the ones currently in
// we delete any that are removed, and add any that are new


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
			validator: (value) => ['prosemirror', 'textarea'].includes(value),
		},
		markdownValue: String,
	},

	data() {
		return {
			internalMarkdown: '',
		}
	},

	mounted() {
		const internalMarkdown = this.internalMarkdown = this.markdownValue

		this.view = new EditorView(this.$refs.prosemirrorEditor, {
			state: this.createState(internalMarkdown),

			dispatchTransaction(transaction) {
				this.view.updateState(this.view.state.apply(transaction))
				this.applyMarkdown(defaultMarkdownSerializer.serialize(this.view.state.doc))
			},
		})

		if (this.mode === 'prosemirror') this.view.focus()
	},

	methods: {
		applyMarkdown(newMarkdown) {
			this.internalMarkdown = newMarkdown
			this.$emit('markdownChange', newMarkdown)
		},

		createState(newMarkdown) {
			return EditorState.create({
				doc: defaultMarkdownParser.parse(newMarkdown),
				plugins: exampleSetup({ schema }),
			})
		},

		checkModeSwitch(currentMode, newMode) {
			if (newMode !== currentMode && newMode === 'prosemirror') {
				const state = this.createState(this.internalMarkdown)
				this.view.updateState(state)
			}
		},

		async acceptFiles(event) {
			foundKeys = new Set()

			const managedImage = await addManagedImage(event.target.files[0])
			const parseInput = `![${managedImage.key}][${managedImage.key}]\n\n[${managedImage.key}]: ${managedImage.url} "${managedImage.name}"`

			defaultMarkdownParser.parse(parseInput)
			console.log(foundKeys)

			managedImages['irrelevant'] = {}

			const toRemove = Object.keys(managedImages).filter(key => !foundKeys.has(key))
			console.log(toRemove)

			// those that manangedImages has but foundKeys doesn't have
			// managedImages / foundKeys = remove
		},

		syncPictures() {
			// imageList
			// diff imageList with a previous?
			// no matter what, we need to parse, because that doesn't happen as a matter of course in either mode
			// once we've parsed, the imageList should be in the correct state
			// and we do the work
		},
	},

	watch: {
		markdownValue(newMarkdownValue) {
			if (newMarkdownValue === this.internalMarkdown) return

			this.internalMarkdown = newMarkdownValue

			this.checkModeSwitch(null, this.mode)
		},

		mode(oldMode, newMode) {
			this.checkModeSwitch(oldMode, newMode)
		}
	}
}


// schema.spec.nodes.update("image", newNodeSpec)

// const defaultMarkdownSerializer = new MarkdownSerializer({
// 	blockquote(state, node) {
// 		state.wrapBlock("> ", null, node, () => state.renderContent(node))
// 	},
// 	code_block(state, node) {
// 		state.write("```" + (node.attrs.params || "") + "\n")
// 		state.text(node.textContent, false)
// 		state.ensureNewLine()
// 		state.write("```")
// 		state.closeBlock(node)
// 	},
// 	heading(state, node) {
// 		state.write(state.repeat("#", node.attrs.level) + " ")
// 		state.renderInline(node)
// 		state.closeBlock(node)
// 	},
// 	horizontal_rule(state, node) {
// 		state.write(node.attrs.markup || "---")
// 		state.closeBlock(node)
// 	},
// 	bullet_list(state, node) {
// 		state.renderList(node, "  ", () => (node.attrs.bullet || "*") + " ")
// 	},
// 	ordered_list(state, node) {
// 		let start = node.attrs.order || 1
// 		let maxW = String(start + node.childCount - 1).length
// 		let space = state.repeat(" ", maxW + 2)
// 		state.renderList(node, space, i => {
// 			let nStr = String(start + i)
// 			return state.repeat(" ", maxW - nStr.length) + nStr + ". "
// 		})
// 	},
// 	list_item(state, node) {
// 		state.renderContent(node)
// 	},
// 	paragraph(state, node) {
// 		state.renderInline(node)
// 		state.closeBlock(node)
// 	},

// 	image(state, node) {
// 		state.write("![" + state.esc(node.attrs.alt || "") + "](" + state.esc(node.attrs.src) +
// 								(node.attrs.title ? " " + state.quote(node.attrs.title) : "") + ")")
// 	},
// 	hard_break(state, node, parent, index) {
// 		for (let i = index + 1; i < parent.childCount; i++)
// 			if (parent.child(i).type != node.type) {
// 				state.write("\\\n")
// 				return
// 			}
// 	},
// 	text(state, node) {
// 		state.text(node.text)
// 	}
// }, {
// 	em: { open: "*", close: "*", mixable: true, expelEnclosingWhitespace: true },
// 	strong: { open: "**", close: "**", mixable: true, expelEnclosingWhitespace: true },
// 	link: {
// 		open: "[",
// 		close(state, mark) {
// 			return "](" + state.esc(mark.attrs.href) + (mark.attrs.title ? " " + state.quote(mark.attrs.title) : "") + ")"
// 		}
// 	},
// 	code: { open: "`", close: "`", escape: false },
// })

// function insertImageItem(nodeType) {
// 	return new MenuItem({
// 		title: "Insert image",
// 		label: "Image",
// 		enable(state) { return canInsert(state, nodeType) },
// 		run(state, _, view) {
// 			let {from, to} = state.selection, attrs = null
// 			if (state.selection instanceof NodeSelection && state.selection.node.type == nodeType)
// 				attrs = state.selection.node.attrs
// 			openPrompt({
// 				title: "Insert image",
// 				fields: {
// 					src: new TextField({label: "Location", required: true, value: attrs && attrs.src}),
// 					title: new TextField({label: "Title", value: attrs && attrs.title}),
// 					alt: new TextField({label: "Description",
// 															value: attrs ? attrs.alt : state.doc.textBetween(from, to, " ")})
// 				},
// 				callback(attrs) {
// 					view.dispatch(view.state.tr.replaceSelectionWith(nodeType.createAndFill(attrs)))
// 					view.focus()
// 				}
// 			})
// 		}
// 	})
// }

// // https://stackoverflow.com/questions/11076975/insert-text-into-textarea-at-cursor-position-javascript
// function insertAtCursor(el, value) {
// 	// IE support
// 	if (document.selection) {
// 		el.focus()
// 		const sel = document.selection.createRange()
// 		sel.text = value
// 	}
// 	// Microsoft Edge, Mozilla, others
// 	else if (el.selectionStart || el.selectionStart == '0') {
// 		const startPos = el.selectionStart
// 		const endPos = el.selectionEnd

// 		el.value = el.value.substring(0, startPos)
// 			+ value
// 			+ el.value.substring(endPos, el.value.length)

// 		const pos = startPos + value.length
// 		el.focus()
// 		el.setSelectionRange(pos, pos)
// 	}
// 	else {
// 		el.value += value
// 	}

// 	const eventType = 'input'
// 	if ('createEvent' in document) {
// 		// modern browsers, IE9+
// 		const e = document.createEvent('HTMLEvents')
// 		e.initEvent(eventType, false, true)
// 		el.dispatchEvent(e)
// 	}
// 	else {
// 		// IE 8
// 		const e = document.createEventObject()
// 		e.eventType = eventType
// 		el.fireEvent('on' + e.eventType, e)
// 	}
// }
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
