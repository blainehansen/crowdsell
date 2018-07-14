import axios from 'axios'
import config from './config'

axios.defaults.baseURL = config.baseURL
axios.defaults.responseType = 'json'
// axios.interceptors.response.use(null, function (error) {
// 	return Promise.reject(error)
// })

export const publicHttp = axios.create()
export const privateHttp = axios.create({ baseURL: config.baseURL + '/secure' })

export const privateApi = {
	uploadFile(url, file) {
		const formData = new FormData()
		formData.append('file', file)
		return privateHttp.post(url, formData, { headers: {'Content-Type': 'multipart/form-data'} })
	},

	uploadProfilePicture(hash, type, file) {
		return this.uploadFile(`/profile-image/${hash}/${type}`, file)
	},

	saveProject: (projectId, projectPatches) => {
		return projectId === null
			? privateHttp.post(`/projects`, projectPatches)
			: privateHttp.patch(`/projects/${projectId}`, projectPatches)
	},
}

export default {
	login: (email, password) => publicHttp.post('/login', { email, password }),
	createUser: (name, email, password) => publicHttp.post('/create-user', { name, email, password }),
	getProjects: () => publicHttp.get('/projects'),

	getProjectById: (projectId) => publicHttp.get(`/projects/${projectId}`)
	getProjectBySlug: (projectSlug) => publicHttp.get(`/projects/${projectId}`)
}
