import axios from 'axios'
import config from '@/config'

import { cloneDeep } from 'lodash'

axios.defaults.responseType = 'json'
// axios.interceptors.response.use(null, function (error) {
// 	return Promise.reject(error)
// })

export const publicHttp = axios.create({ baseURL: config.API_URL })
publicHttp.defaults.headers = cloneDeep(publicHttp.defaults.headers)
export const privateHttp = axios.create({ baseURL: config.API_URL + '/secure' })
privateHttp.defaults.headers = cloneDeep(privateHttp.defaults.headers)



export const publicApi = {
	login: (email, password) => publicHttp.post('/login', { email, password }),
	createUser: (name, email, password) => publicHttp.post('/create-user', { name, email, password }),
	getProjects: () => publicHttp.get('/projects'),

	// getProjectById: (projectId) => publicHttp.get(`/projects/${projectId}`)
	// getProjectBySlug: (projectSlug) => publicHttp.get(`/projects/${projectId}`)
}

export const privateApi = {
	fetchProfileUploadSignature: () => privateHttp.post('/user/profile-image/sign'),
	confirmProfileUpload: (signature, timestamp, version) => privateHttp.post(`/user/profile-image/confirm`, { signature, timestamp, version }),

	fetchProjectUploadSignatures: (projectId, fileHashes) => privateHttp.post(`/project/${projectId}/uploads/sign`, { hashes: fileHashes }),
	confirmProjectUploads: (projectId, confirmationPayloads) => privateHttp.post(`/project/${projectId}/uploads/confirm`, { confirmations: confirmationPayloads }),

	fetchFullUser: () => privateHttp.get('/user'),
	changeSlug: (newSlug) => privateHttp.put('/user/slug', { slug: newSlug }),
	changePassword: (oldPassword, newPassword) => privateHttp.put('/user/password', { oldPassword, newPassword }),
	saveUser: (userPatches) => privateHttp.patch('/user', userPatches),

	saveProject(projectId, projectPatches) {
		return projectId === null
			? privateHttp.post(`/projects`, projectPatches)
			: privateHttp.patch(`/projects/${projectId}`, projectPatches)
	},

	generateCardToken: () => privateHttp.post('/user/card-token'),
	generateBankToken: () => privateHttp.post('/user/bank-token'),
}


const imagesHttp = axios.create({
	baseURL: config.CDN_API_ENDPOINT,
	withCredentials: false,
	headers: {},
})
imagesHttp.defaults.headers = cloneDeep(imagesHttp.defaults.headers)

export const imagesApi = {
	async postFile(fileDataPromise, route, signature, objectName, timestamp, preset, requestConfig = {}) {
		const formData = new FormData()
		formData.append('api_key', config.CDN_API_KEY)
		formData.append('file', fileDataPromise)
		formData.append('signature', signature)
		formData.append('public_id', objectName)
		formData.append('timestamp', timestamp)

		if (preset) formData.append('upload_preset', preset)

		return imagesHttp.post(route, formData, requestConfig)
	},

	async uploadFile(file, route, signature, objectName, timestamp, preset, requestConfig = undefined) {
		const chunkLimit = 6291456
		const fileSize = file.size

		if (fileSize > chunkLimit) {
			const slicePromises = []
			const uploadId = randomId()
			requestConfig = requestConfig || {}

			const createSlicePromise = (startIndex, endIndex = undefined) => {
				const bytesEndIndex = (endIndex || fileSize) - 1
				const sliceEndIndex = endIndex

				return this.postFile(
					file.slice(startIndex, sliceEndIndex),
					route, signature, objectName, timestamp, preset,
					{
						headers: {
							'Content-Type': 'multipart/form-data',
							'X-Unique-Upload-Id': uploadId,
							'Content-Range': `bytes ${startIndex}-${bytesEndIndex}/${fileSize}`,
							...requestConfig,
						},
						onUploadProgress: (progressEvent) => {
							if (progressEvent.lengthComputable) {
								console.log((progressEvent.loaded / progressEvent.total) * 100)
							}
						}
					},
				)
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

		return this.postFile(file, route, signature, objectName, timestamp, preset, requestConfig)
	},

	async uploadProfileImage(file) {
		const { data: { signature, objectName, timestamp } } = await privateApi.fetchProfileUploadSignature(signature, objectName, timestamp)

		const response = await this.uploadFile(
			file,
			config.CDN_API_IMAGES_ROUTE,
			signature,
			objectName,
			timestamp,
			config.CDN_API_PROFILE_IMAGES_PRESET,
		)

		await privateApi.confirmProfileUpload(signature, timestamp, version.toString())

		return version
	},

	async uploadProjectImage(file, objectName, signature, timestamp, progressFunction) {
		const { data: { version } } = await this.uploadFile(
			file,
			config.CDN_API_IMAGES_ROUTE,
			signature,
			objectName,
			timestamp,
			config.CDN_API_PROJECT_IMAGES_PRESET,
			{ onUploadProgress: progressFunction },
		)

		return version
	},

	async uploadVideo(file) {
		const { data: { signature, objectName, timestamp } } = await privateApi.fetchProfileUploadSignature()

		const response = await this.uploadFile(
			file,
			'/raw/upload',
			signature,
			objectName,
			timestamp,
		)

		return response
	},

	// deleteImage(objectName, signature, timestamp) {
	// 	return imagesHttp.post(config.CDN_API_IMAGES_DELETE_ROUTE, {
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
