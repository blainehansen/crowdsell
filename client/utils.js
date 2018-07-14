import xxh from 'xxhashjs'

export async function sampleHashFile(file) {
	return new Promise(function (resolve, reject) {
		const reader = new FileReader()
		const hasher = xxh.h64(0xABCD)

		const oneThird = Math.floor(file.size / 3)
		const twoThird = Math.floor(file.size * 2 / 3)
		const slices = [
			file.slice(0, 256),
			file.slice(oneThird, oneThird + 256),
			file.slice(twoThird, twoThird + 256),
			file.slice(-256),
		]

		reader.onloadend = (event) => {
			if ( event.target.readyState !== FileReader.DONE ) return
			hasher.update(event.target.result)

			if (slices.length > 0) nextSlice()
			else resolve(hasher.digest().toString(36))
		}

		function nextSlice() {
			const slice = slices.pop()
			reader.readAsBinaryString(slice)
		}
		nextSlice()
	})
}

export function formatSpacesUrl(inputUrl) {
	return `https://blaine-final-spaces-test.nyc3.cdn.digitaloceanspaces.com/${inputUrl}`
}
