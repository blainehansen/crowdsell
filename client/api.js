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


export const privateApi = {
	// uploadFile(url, file) {
	// 	const formData = new FormData()
	// 	formData.append('file', file)
	// 	return privateHttp.post(url, formData, { headers: { 'Content-Type': 'multipart/form-data' } })
	// },

	// uploadProfilePicture(hash, type, file) {
	// 	return this.uploadFile(`/user/profile-image/${hash}/${type}`, file)
	// },

	fetchUploadSignature: () => privateHttp.post(`/user/profile-image/sign`),
	confirmUpload: (signature, timestamp, version) => privateHttp.post(`/user/profile-image/confirm`, { signature, timestamp, version }),

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

export const publicApi = {
	login: (email, password) => publicHttp.post('/login', { email, password }),
	createUser: (name, email, password) => publicHttp.post('/create-user', { name, email, password }),
	getProjects: () => publicHttp.get('/projects'),

	// getProjectById: (projectId) => publicHttp.get(`/projects/${projectId}`)
	// getProjectBySlug: (projectSlug) => publicHttp.get(`/projects/${projectId}`)
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
		const bufferPromise = getFileData(file)

		const { data: { objectName, signature, timestamp } } = await privateApi.fetchUploadSignature()

		const { data: { version } } = await imagesHttp.post(config.CDN_API_IMAGES_ROUTE, {
			file: await bufferPromise,
			timestamp,
			api_key: config.CDN_API_KEY,
			public_id: objectName,
			signature,
			upload_preset: config.CDN_API_PROFILE_IMAGES_PRESET,
		})

		await privateApi.confirmUpload(signature, timestamp, version.toString())

		return version
	}
}
