import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

const ApiService = {
	init() {
		Vue.use(VueAxios, axios)
		Vue.axios.defaults.baseURL = 'http://localhost:5555/api'
	},

	query(resource, params) {
		return Vue.axios.get(resource, params).catch((error) => {
			throw new Error(`[StayUp] ApiService ${error}`)
		})
	},

	get(resource) {
		return Vue.axios.get(`${resource}`).catch((error) => {
			throw new Error(`[StayUp] ApiService ${error}`)
		})
	},

	post(resource, data) {
		return Vue.axios.post(`${resource}`, data)
	},

	update(resource, slug, params) {
		return Vue.axios.put(`${resource}/${slug}`, params)
	},

	put(resource, params) {
		return Vue.axios.put(`${resource}`, params)
	},

	delete(resource) {
		return Vue.axios.delete(resource).catch((error) => {
			throw new Error(`[StayUp] ApiService ${error}`)
		})
	}
}

export const SvcService = {
	get() {
		return ApiService.get('ws')
	},

	post(data) {
		return ApiService.post('service', data)
	}
}

export default ApiService
