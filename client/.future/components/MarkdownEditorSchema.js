import { Schema } from 'prosemirror-model'

export const managedImagePrefix = 'managed-image:'

export const schema = new Schema({
	nodes: {
		doc: {
			content: "block+"
		},

		paragraph: {
			content: "inline*",
			group: "block",
			parseDOM: [{tag: "p"}],
			toDOM() { return ["p", 0] }
		},

		blockquote: {
			content: "block+",
			group: "block",
			parseDOM: [{tag: "blockquote"}],
			toDOM() { return ["blockquote", 0] }
		},

		horizontal_rule: {
			group: "block",
			parseDOM: [{tag: "hr"}],
			toDOM() { return ["div", ["hr"]] }
		},

		heading: {
			attrs: {level: {default: 1}},
			content: "inline*",
			group: "block",
			defining: true,
			parseDOM: [
				{tag: "h1", attrs: {level: 1}},
				{tag: "h2", attrs: {level: 2}},
				{tag: "h3", attrs: {level: 3}},
				{tag: "h4", attrs: {level: 4}},
				{tag: "h5", attrs: {level: 5}},
				{tag: "h6", attrs: {level: 6}}
			],
			toDOM(node) { return ["h" + node.attrs.level, 0] }
		},

		code_block: {
			content: "text*",
			group: "block",
			code: true,
			defining: true,
			attrs: {params: {default: ""}},
			parseDOM: [{tag: "pre", preserveWhitespace: true, getAttrs: node => (
				{params: node.getAttribute("data-params") || ""}
			)}],
			toDOM(node) { return ["pre", node.attrs.params ? {"data-params": node.attrs.params} : {}, ["code", 0]] }
		},

		ordered_list: {
			content: "list_item+",
			group: "block",
			attrs: {order: {default: 1}, tight: {default: false}},
			parseDOM: [{tag: "ol", getAttrs(dom) {
				return {order: dom.hasAttribute("start") ? +dom.getAttribute("start") : 1, tight: dom.hasAttribute("data-tight")}
			}}],
			toDOM(node) {
				return ["ol", {start: node.attrs.order == 1 ? null : node.attrs.order, "data-tight": node.attrs.tight ? "true" : null}, 0]
			}
		},

		bullet_list: {
			content: "list_item+",
			group: "block",
			attrs: {tight: {default: false}},
			parseDOM: [{tag: "ul", getAttrs: dom => ({tight: dom.hasAttribute("data-tight")})}],
			toDOM(node) { return ["ul", {"data-tight": node.attrs.tight ? "true" : null}, 0] }
		},

		list_item: {
			content: "paragraph block*",
			defining: true,
			parseDOM: [{tag: "li"}],
			toDOM() { return ["li", 0] }
		},

		text: {
			group: "inline",
			toDOM(node) { return node.text }
		},

		image: {
			inline: true,
			attrs: {
				key: { default: null },
				src: {},
				alt: { default: null },
				title: { default: null }
			},
			group: "inline",
			draggable: true,
			parseDOM: [{tag: "img[src]", getAttrs(dom) {
				return {
					key: null,
					src: dom.getAttribute("src"),
					title: dom.getAttribute("title"),
					alt: dom.getAttribute("alt")
				}
			}}],
			toDOM(node) {
				const attrs = node.attrs
				return ["img", { src: attrs.src, title: attrs.title, alt: attrs.alt }]
			},
		},

		hard_break: {
			inline: true,
			group: "inline",
			selectable: false,
			parseDOM: [{tag: "br"}],
			toDOM() { return ["br"] }
		}
	},

	marks: {
		em: {
			parseDOM: [{tag: "i"}, {tag: "em"}, {style: "font-style", getAttrs: value => value == "italic" && null}],
			toDOM() { return ["em"] }
		},

		strong: {
			parseDOM: [{tag: "b"}, {tag: "strong"}, {style: "font-weight", getAttrs: value => /^(bold(er)?|[5-9]\d{2,})$/.test(value) && null}],
			toDOM() { return ["strong"] }
		},

		link: {
			attrs: {
				href: {},
				title: {default: null}
			},
			inclusive: false,
			parseDOM: [{tag: "a[href]", getAttrs(dom) {
				return {href: dom.getAttribute("href"), title: dom.getAttribute("title")}
			}}],
			toDOM(node) { return ["a", node.attrs] }
		},

		code: {
			parseDOM: [{tag: "code"}],
			toDOM() { return ["code"] }
		}
	}
})


export const baseMarkdownParserArgs = {
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
		// getAttrs
	},

	hardbreak: { node: "hard_break" },

	em: { mark: "em" },
	strong: { mark: "strong" },
	link: { mark: "link", getAttrs: tok => ({ href: tok.attrGet("href"), title: tok.attrGet("title") || null }) },
	code_inline: { mark: "code" },
}


export const markdownSerializerArgs = [{
	blockquote(state, node) {
		state.wrapBlock("> ", null, node, () => state.renderContent(node))
	},
	code_block(state, node) {
		state.write("```" + (node.attrs.params || "") + "\n")
		state.text(node.textContent, false)
		state.ensureNewLine()
		state.write("```")
		state.closeBlock(node)
	},
	heading(state, node) {
		state.write(state.repeat("#", node.attrs.level) + " ")
		state.renderInline(node)
		state.closeBlock(node)
	},
	horizontal_rule(state, node) {
		state.write(node.attrs.markup || "---")
		state.closeBlock(node)
	},
	bullet_list(state, node) {
		state.renderList(node, "  ", () => (node.attrs.bullet || "*") + " ")
	},
	ordered_list(state, node) {
		let start = node.attrs.order || 1
		let maxW = String(start + node.childCount - 1).length
		let space = state.repeat(" ", maxW + 2)
		state.renderList(node, space, i => {
			let nStr = String(start + i)
			return state.repeat(" ", maxW - nStr.length) + nStr + ". "
		})
	},
	list_item(state, node) {
		state.renderContent(node)
	},
	paragraph(state, node) {
		state.renderInline(node)
		state.closeBlock(node)
	},

	image(state, node) {
		const alt = state.esc(node.attrs.alt || "")
		const src = state.esc(node.attrs.key || node.attrs.src)
		const title = node.attrs.title ? " " + state.quote(node.attrs.title) : ""
		state.write(`![${alt}](${src}${title})`)
	},
	hard_break(state, node, parent, index) {
		for (let i = index + 1; i < parent.childCount; i++)
			if (parent.child(i).type != node.type) {
				state.write("\\\n")
				return
			}
	},
	text(state, node) {
		state.text(node.text)
	}
}, {
	em: { open: "*", close: "*", mixable: true, expelEnclosingWhitespace: true },
	strong: { open: "**", close: "**", mixable: true, expelEnclosingWhitespace: true },
	link: {
		open: "[",
		close(state, mark) {
			return "](" + state.esc(mark.attrs.href) + (mark.attrs.title ? " " + state.quote(mark.attrs.title) : "") + ")"
		}
	},
	code: { open: "`", close: "`", escape: false },
}]
