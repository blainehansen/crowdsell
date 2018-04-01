import http from 'axios'

export default {
	getProjects: () => http.get('/projects')
}
