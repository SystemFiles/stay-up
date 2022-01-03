const state = {
	tags          : [],
	articles      : [],
	isLoading     : true,
	articlesCount : 0
}

const getters = {}

/* eslint no-param-reassign: ["error", { "props": false }] */
const mutations = {
	[FETCH_START](state) {
		state.isLoading = true
	},
	[FETCH_END](state, { articles, articlesCount }) {
		state.articles = articles
		state.articlesCount = articlesCount
		state.isLoading = false
	},
	[SET_TAGS](state, tags) {
		state.tags = tags
	},
	[UPDATE_ARTICLE_IN_LIST](state, data) {
		state.articles = state.articles.map((article) => {
			if (article.slug !== data.slug) {
				return article
			}
			// We could just return data, but it seems dangerous to
			// mix the results of different api calls, so we
			// protect ourselves by copying the information.
			article.favorited = data.favorited
			article.favoritesCount = data.favoritesCount
			return article
		})
	}
}

export default {
	state,
	getters,
	actions,
	mutations
}
