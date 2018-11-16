import axios from 'axios'

axios.defaults.responseType = 'json'
// axios.interceptors.response.use(null, function (error) {
// 	return Promise.reject(error)
// })

const API_URL = process.env.API_URL
export const publicGolangHttp = axios.create({ baseURL: API_URL, headers: {} })
export const secureGolangHttp = axios.create({ baseURL: API_URL + '/secure', headers: {} })

const GQL_API_URL = process.env.GQL_API_URL
export const publicGqlHttp = axios.create({ baseURL: GQL_API_URL, headers: {} })
export const secureGqlHttp = axios.create({ baseURL: GQL_API_URL + '/secure', headers: {} })

import publicQueries from '@/queries/public-queries.gql'
// import secureQueries from '@/queries/secure-queries.gql'
// import secureMutations from '@/queries/secure-mutations.gql'


export const publicApi = {
	login: (email, password) => publicGolangHttp.post('/login', { email, password }),
	createUser: (name, email, password) => publicGolangHttp.post('/create-user', { name, email, password }),
	getPeople: () => publicGqlHttp.get(publicQueries.firstPeople)

	// getProjectById: (projectId) => publicGolangHttp.get(`/projects/${projectId}`)
	// getProjectBySlug: (projectSlug) => publicGolangHttp.get(`/projects/${projectId}`)
}

export const secureApi = {
	fetchProfileUploadSignature: () => secureGolangHttp.post('/user/profile-image/sign'),
	confirmProfileUpload: (signature, timestamp, version) => secureGolangHttp.post(`/user/profile-image/confirm`, { signature, timestamp, version }),

	fetchProjectUploadSignatures: (projectId, fileHashes) => secureGolangHttp.post(`/project/${projectId}/uploads/sign`, { hashes: fileHashes }),
	confirmProjectUploads: (projectId, confirmationPayloads) => secureGolangHttp.post(`/project/${projectId}/uploads/confirm`, { confirmations: confirmationPayloads }),

	fetchFullUser: () => secureGolangHttp.get('/user'),
	changeSlug: (newSlug) => secureGolangHttp.put('/user/slug', { slug: newSlug }),
	changePassword: (oldPassword, newPassword) => secureGolangHttp.put('/user/password', { oldPassword, newPassword }),
	saveUser: (userPatches) => secureGolangHttp.patch('/user', userPatches),

	saveProject(projectId, projectPatches) {
		return projectId === null
			? secureGolangHttp.post(`/projects`, projectPatches)
			: secureGolangHttp.patch(`/projects/${projectId}`, projectPatches)
	},

	generateCardToken: () => secureGolangHttp.post('/user/card-token'),
	generateBankToken: () => secureGolangHttp.post('/user/bank-token'),
}


const imagesHttp = axios.create({
	baseURL: process.env.CDN_API_ENDPOINT,
	withCredentials: false,
	headers: {},
})

async function postFile(fileSlice, route, signature, objectName, timestamp, preset, requestConfig = {}) {
	const formData = new FormData()
	formData.append('api_key', process.env.CDN_API_KEY)
	formData.append('file', fileSlice)
	formData.append('signature', signature)
	formData.append('public_id', objectName)
	formData.append('timestamp', timestamp)

	if (preset) formData.append('upload_preset', preset)

	return imagesHttp.post(route, formData, requestConfig)
}

async function uploadFile(file, route, signature, objectName, timestamp, preset, requestConfig = {}) {
	const chunkLimit = 6291456
	const fileSize = file.size

	if (fileSize > chunkLimit) {
		const slicePromises = []
		const uploadId = randomId()

		function createSlicePromise(startIndex, endIndex = undefined) {
			const bytesEndIndex = (endIndex || fileSize) - 1
			const sliceEndIndex = endIndex

			const requestConfig = {
				headers: {
					'Content-Type': 'multipart/form-data',
					'X-Unique-Upload-Id': uploadId,
					'Content-Range': `bytes ${startIndex}-${bytesEndIndex}/${fileSize}`,
					...requestConfig,
				},
				onUploadProgress(progressEvent) {
					if (progressEvent.lengthComputable) {
						console.log((progressEvent.loaded / progressEvent.total) * 100)
					}
				}
			}

			return postFile(file.slice(startIndex, sliceEndIndex), route, signature, objectName, timestamp, preset, requestConfig)
		}

		let currentIndex = 0
		// file size is non-inclusive
		while (currentIndex + chunkLimit < fileSize) {
			const startIndex = currentIndex
			const endIndex = currentIndex + chunkLimit
			slicePromises.push(createSlicePromise(startIndex, endIndex))

			currentIndex = endIndex
		}

		// upload the last one
		slicePromises.push(createSlicePromise(currentIndex))

		const results = await Promise.all(slicePromises)
		for (const { data } of results) {
			if (data.done) return data.version
		}
		throw new Error("no version created")
	}

	return postFile(file, route, signature, objectName, timestamp, preset, requestConfig)
}

export const imagesApi = {
	// async uploadProfileImage(file) {
	// 	const { data: { signature, objectName, timestamp } } = await secureApi.fetchProfileUploadSignature(signature, objectName, timestamp)

	// 	const response = await uploadFile(
	// 		file,
	// 		process.env.CDN_API_IMAGES_ROUTE,
	// 		signature,
	// 		objectName,
	// 		timestamp,
	// 		process.env.CDN_API_PROFILE_IMAGES_PRESET,
	// 	)

	// 	await secureApi.confirmProfileUpload(signature, timestamp, version.toString())

	// 	return version
	// },

	async uploadProjectImage(file, objectName, signature, timestamp, progressFunction) {
		const { data: { version } } = await uploadFile(
			file,
			process.env.CDN_API_IMAGES_ROUTE,
			signature,
			objectName,
			timestamp,
			process.env.CDN_API_PROJECT_IMAGES_PRESET,
			{ onUploadProgress: progressFunction },
		)

		return version
	},

	async uploadVideo(file) {
		const { data: { signature, objectName, timestamp } } = await secureApi.fetchProfileUploadSignature()

		const response = await uploadFile(
			file,
			'/raw/upload',
			signature,
			objectName,
			timestamp,
		)

		return response
	},

	// deleteImage(objectName, signature, timestamp) {
	// 	return imagesHttp.post(process.env.CDN_API_IMAGES_DELETE_ROUTE, {
	// 		public_id: objectName,
	// 		signature,
	// 		timestamp,
	// 		invalidate: true,
	// 	})
	// },
}



function randomId(chunks = 6) {
	let id = ""
	for (let i = 0; i < chunks; i++) {
		id += Math.random().toString(36).substring(2, 15)
	}

	return id
}
