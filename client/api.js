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
	// uploadFile(url, file) {
	// 	const formData = new FormData()
	// 	formData.append('file', file)
	// 	return privateHttp.post(url, formData, { headers: { 'Content-Type': 'multipart/form-data' } })
	// },

	// uploadProfilePicture(hash, type, file) {
	// 	return this.uploadFile(`/user/profile-image/${hash}/${type}`, file)
	// },

	fetchProfileUploadSignature: () => privateHttp.post('/user/profile-image/sign'),
	confirmProfileUpload: (signature, timestamp, version) => privateHttp.post(`/user/profile-image/confirm`, { signature, timestamp, version }),

	fetchProjectUploadSignatures: (projectId, fileHashes) => privateHttp.post(`/project/${projectId}/uploads/sign`, { hashes: fileHashes }),
	comfirmProjectUploads: (projectId, confirmationPayloads) => privateHttp.post(`/project/${projectId}/uploads/confirm`, { confirmations: confirmationPayloads }),

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

import { getFileData } from '@/utils'

export const imagesApi = {
	async uploadProfileImage(file) {
		const fileDataPromise = getFileData(file)

		const { data: { objectName, signature, timestamp } } = await privateApi.fetchProfileUploadSignature()

		const { data: { version } } = await imagesHttp.post(config.CDN_API_IMAGES_ROUTE, {
			file: await fileDataPromise,
			timestamp,
			api_key: config.CDN_API_KEY,
			public_id: objectName,
			signature,
			upload_preset: config.CDN_API_PROFILE_IMAGES_PRESET,
		})

		await privateApi.confirmProfileUpload(signature, timestamp, version.toString())

		return version
	},

	async uploadProjectImage(file, objectName, signature, timestamp, progressFunction) {
		const fileDataPromise = getFileData(file)

		const { data: { version } } = await imagesHttp.post(config.CDN_API_IMAGES_ROUTE, {
			file: await fileDataPromise,
			timestamp,
			api_key: config.CDN_API_KEY,
			public_id: objectName,
			signature,
			upload_preset: config.CDN_API_PROJECT_IMAGES_PRESET,
		}, {
			onUploadProgress: progressFunction,
		})

		return version
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
