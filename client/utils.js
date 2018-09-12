import xxh from 'xxhashjs'

export function sampleHashFile(file) {
	return new Promise(function (resolve, reject) {
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

export function getFileData(file) {
	return new Promise(function (resolve, reject) {
		const reader = new FileReader()

		reader.readAsDataURL(file)

		reader.onloadend = (event) => {
			if (event.target.readyState !== FileReader.DONE)
				return

			resolve(event.target.result)
		}
	})
}

import config from '@/config'
import store from '@/vuex'

export function formatSpacesUrl(inputUrl) {
	return `https://${config.SPACES_BUCKET_NAME}.${config.CDN_BASENAME}/${inputUrl}`
}


export function formatProfileImageUrl(version) {
	const userId = store.getters['auth/userId']

	return `${config.CDN_ENDPOINT}${config.CDN_IMAGES_ROUTE}/v${version}/profile-images/${userId}.png`
}
