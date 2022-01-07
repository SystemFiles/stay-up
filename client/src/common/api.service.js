import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { API_BASE_URL } from '../config'

const ApiService = {
	init() {
		Vue.use(VueAxios, axios)
		Vue.axios.defaults.baseURL = API_BASE_URL
	},

	get(resource) {
		return Vue.axios.get(`${resource}`).catch((error) => {
			throw new Error(`[StayUp] ApiService ${error}`)
		})
	},

	post(resource, data) {
		return Vue.axios.post(`${resource}`, data)
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
	getWebsocket() {
		return ApiService.get('ws')
	},

	get(id) {
		return ApiService.get(`service/${id}`)
	},

	post(data) {
		return ApiService.post('service', data)
	},

	put(attr, val) {
		return ApiService.put('service', {
			attribute : attr,
			new_value : val
		})
	},

	delete(id) {
		return ApiService.delete(`service/${id}`)
	}
}

export default ApiService
