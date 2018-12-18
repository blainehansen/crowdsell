import xxh from 'xxhashjs'

export function sampleHashFile(file) {
	return new Promise(function (resolve, _reject) {
		const reader = new FileReader()
		const hasher = xxh.h64(0xABCD)

		const oneThird = Math.floor(file.size / 3)
		const twoThird = Math.floor(file.size * 2 / 3)
		const sliceSize = 512
		const slices = [
			file.slice(0, sliceSize),
			file.slice(oneThird, oneThird + sliceSize),
			file.slice(twoThird, twoThird + sliceSize),
			file.slice(-sliceSize),
		]

		reader.onloadend = (event) => {
			if ( event.target.readyState !== FileReader.DONE )
				return
			hasher.update(event.target.result)

			if (slices.length > 0)
				nextSlice()
			else
				resolve(hasher.digest().toString(36))
		}

		function nextSlice() {
			const slice = slices.pop()
			reader.readAsBinaryString(slice)
		}
		nextSlice()
	})
}

// https://stackoverflow.com/questions/11076975/insert-text-into-textarea-at-cursor-position-javascript
export function insertToElement(el, cursorValue = undefined, endValue = undefined) {
	if (cursorValue) {
		// IE support
		if (document.selection) {
			el.focus()
			const sel = document.selection.createRange()
			sel.text = cursorValue
		}
		// Microsoft Edge, Mozilla, others
		else if (el.selectionStart || el.selectionStart == '0') {
			const startPos = el.selectionStart
			const endPos = el.selectionEnd

			el.value = el.value.substring(0, startPos)
				+ cursorValue
				+ el.value.substring(endPos, el.value.length)

			const pos = startPos + cursorValue.length
			el.focus()
			el.setSelectionRange(pos, pos)
		}
		else {
			el.value += cursorValue
		}
	}
	if (endValue) {
		el.value += endValue
	}

	const eventType = 'input'
	if ('createEvent' in document) {
		// modern browsers, IE9+
		const e = document.createEvent('HTMLEvents')
		e.initEvent(eventType, false, true)
		el.dispatchEvent(e)
	}
	else {
		// IE 8
		const e = document.createEventObject()
		e.eventType = eventType
		el.fireEvent('on' + e.eventType, e)
	}
}


import store from '@/vuex'

export function formatSpacesUrl(inputUrl) {
	return `https://${process.env.SPACES_BUCKET_NAME}.${process.env.CDN_BASENAME}/${inputUrl}`
}


export function formatProfileImageUrl(version) {
	const userId = store.getters['auth/userId']

	return `${process.env.CDN_ENDPOINT}${process.env.CDN_IMAGES_ROUTE}/v${version}/profile-images/${userId}.png`
}
