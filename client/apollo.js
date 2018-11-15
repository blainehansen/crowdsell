import Vue from 'vue'
import VueApollo from 'vue-apollo'

import { ApolloClient } from 'apollo-client'

import { HttpLink } from 'apollo-link-http'
import { ApolloLink, split } from 'apollo-link'

import { InMemoryCache } from 'apollo-cache-inmemory'

const publicQueryMap = require('@/queries/public-queries.json')
const secureQueryMap = require('@/queries/secure-queries.json')
const mutationMap = require('@/queries/secure-mutations.json')

import querystring from 'querystring'

function customFetch(uri, options) {
	const [urlBase, urlQuery] = uri.split('?')

	const variablesString = querystring.parse(urlQuery).variables
	const variablesObj = JSON.parse(variablesString)

	const finalUri = urlBase + '?' + querystring.stringify(variablesObj)
	return fetch(finalUri, options)
}

// split(
// 	function(operation) {

// 	},
// 	// authenticated
// 	// unauthenticated
// )

import store from './vuex'

const authLink = new ApolloLink((operation, forward) => {
  operation.setContext(context => {
  	const token = store.state.auth.token
  	if (token) {
	  	context.headers = context.headers || {}
	  	context.headers.authorization = token
  	}
  	return context
  })
  return forward(operation)
})


const link = ApolloLink.from([
	new ApolloLink((operation, forward) => {
		const queryDocument = operation.query.loc.source.body
		const queryHash = queryMap[queryDocument]
		operation.variables.$h = queryHash

		operation.setContext({
			http: {
				includeQuery: false,
			}
		})

		return forward(operation)
	}),

	new HttpLink({
		// You should use an absolute URL here
		uri: 'http://localhost:5555/graphql',
		useGETForQueries: true,
		fetch: customFetch,
	}),
])


// Create the apollo client
const apolloClient = new ApolloClient({
	link,
	cache: new InMemoryCache({
		addTypename: false,
	}),
	// connectToDevTools: true,
})

export default const apolloProvider = new VueApollo({
	defaultClient: apolloClient,
})

// Install the vue plugin
Vue.use(VueApollo)
