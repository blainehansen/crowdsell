<template lang="pug">

#home
	h1 Welcome

	div(v-for="person in people")
		p {{ person.slug }}
		p {{ person.name }}

	//- MarkdownEditor(
	//- 	initialMarkdown="**hello**",
	//- 	mode="all",
	//- 	:textareaConfig="{name: 'my-textarea'}",
	//- 	@contentChangeMarkdown="markdownChangedHandler",
	//- )

	//- .upload-tray
	//- 	.everything-good(v-if="allSuccessful") Everything's good!
	//- 	.upload-item(v-for="upload in uploads")
	//- 		.loading(v-if="upload.loading") loading
	//- 		.stuff {{ upload.localUrl }}
	//- 		.stuff(v-if="upload.hashName") {{ upload.hashName }}

	//- 		template(v-else)
	//- 			| {{ upload.remoteVersion }}

	//- 	.add
	//- 		input(type="file", accept="image/png, image/jpeg", multiple, @change="acceptFiles")

	//- .test
	//- 	input(type="file", @change="acceptFiles")



	//- button(@click="sendCardInfo") send
	//- p(v-if="cardInfoSuccess") {{ cardInfoSuccess }}
	//- p(v-if="cardInfoError") {{ cardInfoError }}

</template>

<script>
// import MarkdownEditor from '@/components/MarkdownEditor'

// import { privateApi, publicApi, imagesApi } from '@/api'
import { publicApi } from '@/api'
// import { sampleHashFile } from '@/utils'

export default {
	name: 'home',

	asyncData: {
		people: {
			get: publicApi.getPeople,
			transform: (response) => response.data.people.nodes,
		}
	},

	// components: {
	// 	MarkdownEditor,
	// },

	// data() {
	// 	return {
	// 		uploads: [],
	// 		allSuccessful: false,
	// 	}
	// },

	// methods: {
	// 	async acceptFiles(event) {
	// 		const eventFile = event.target.files[0]
	// 		console.log(eventFile)

	// 		const version = await imagesApi.uploadVideo(eventFile)

	// 		console.log(version)
	// 	},

	// 	markdownChangedHandler(...args) {
	// 		console.log(args)
	// 	},

	// 	async acceptFiles(event) {
	// 		const eventFiles = event.target.files
	// 		const fileHashes = []
	// 		for (let i = 0; i < eventFiles.length; i++) {
	// 			const eventFile = eventFiles[i]
	// 			fileHashes.push(sampleHashFile(eventFile))

	// 			const uploadObject = {
	// 				name: eventFile.name,
	// 				localUrl: URL.createObjectURL(eventFile),

	// 				loading: false,
	// 				progress: 0,
	// 				hashName: null,
	// 				remoteVersion: null,
	// 			}
	// 			this.uploads.push(uploadObject)
	// 		}

	// 		// const projectId = this.projectId
	// 		const projectId = "ZNWGovPn"

	// 		const finishedHashes = await Promise.all(fileHashes)
	// 		const { data: signaturesTimestamps } = await privateApi.fetchProjectUploadSignatures(projectId, finishedHashes)

	// 		const allUploads = []
	// 		for (var i = 0; i < eventFiles.length; i++) {
	// 			const uploadObject = this.uploads[i]
	// 			const eventFile = eventFiles[i]
	// 			const { objectName, signature, timestamp } = signaturesTimestamps[i]
	// 			const hash = finishedHashes[i]

	// 			uploadObject.hashName = hash

	// 			const progressFunction = (progressEvent) => {
	// 				if (progressEvent.lengthComputable) {
	// 					uploadObject.progress = (progressEvent.loaded / progressEvent.total) * 100
	// 				}
	// 			}
	// 			uploadObject.loading = true
	// 			const imagePromise = imagesApi.uploadProjectImage(eventFile, objectName, signature, timestamp, progressFunction)
	// 				.then((version) => {
	// 					version = version.toString()
	// 					uploadObject.remoteVersion = version

	// 					uploadObject.loading = false

	// 					return { version, signature, timestamp, hash }
	// 				})

	// 			allUploads.push(imagePromise)
	// 		}

	// 		const finishedUploads = await Promise.all(allUploads)
	// 		await privateApi.confirmProjectUploads(projectId, finishedUploads)
	// 		this.allSuccessful = true
	// 	},
	// },

	// data() {
	// 	return {
	// 		cardInfoSuccess: null,
	// 		cardInfoError: null,
	// 	}
	// },
	// methods: {
	// 	async sendCardInfo() {
	// 		const { data: cardToken } = await privateApi.generateCardToken()
	// 		console.log(cardToken)

	// 		promisepay.createCardAccount(cardToken, {
	// 			full_name: "Bella Buyer",
	// 			number: "4111111111111111",
	// 			expiry_month: "02",
	// 			expiry_year: "2022",
	// 			cvv: "123"
	// 		}, (data) => {
	// 			this.cardInfoSuccess = data
	// 		}, (data) => {
	// 			this.cardInfoError = data
	// 		})
	// 	}
	// }
}
</script>

<style lang="sass">
</style>
