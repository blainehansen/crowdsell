import axios from 'axios'
import config from './config'

axios.defaults.baseURL = config.baseURL
axios.defaults.responseType = 'json'
// axios.defaults.headers.post['Content-Type'] = 'application/json'
// axios.interceptors.response.use(null, function (error) {
// 	return Promise.reject(error)
// })

export const publicHttp = axios.create()
export const privateHttp = axios.create({ baseURL: config.baseURL + '/secure' })

export default {
	getProjects: () => publicHttp.get('/projects'),

	login: (email, password) => publicHttp.post('/login', { email, password }),
	createUser: (name, email, password) => publicHttp.post('/create-user', { name, email, password }),
}
