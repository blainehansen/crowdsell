const { parse: parseGraphql, execute: executeGraphql } = require('graphql')

function reduceQueryMap(intitialQueryMap) {
	const queryMap = {}

	for (const [query, hash] of Object.entries(intitialQueryMap)) {
		queryMap[hash] = parseGraphql(query)
	}

	return queryMap
}


const fs = require('fs')
const publicQueryMap = reduceQueryMap(JSON.parse(fs.readFileSync('/public-queries.json')))
const secureQueryMap = reduceQueryMap(JSON.parse(fs.readFileSync('/secure-queries.json')))
const secureMutationMap = reduceQueryMap(JSON.parse(fs.readFileSync('/secure-mutations.json')))

// const environment = require('./environment.js')

// const connectionString = "postgres://postgraphile_user:postgraphile-password@database:5432/dev_database"
const connectionString = "postgres://user:asdf@database:5432/dev_database"

// Create a postgres pool for efficiency
const pg = require("pg")
const pgPool = new pg.Pool({ connectionString })

let graphqlSchema


async function main() {
	// const querystring = require('querystring')

	// creates postgraphile schema
	const { createPostGraphileSchema } = require('postgraphile-core')
	const PgSimplifyInflectorPlugin = require("@graphile-contrib/pg-simplify-inflector")


	graphqlSchema = await createPostGraphileSchema(
		connectionString,
		["public"],
		{
			dynamicJson: true,
			appendPlugins: [PgSimplifyInflectorPlugin],
			ignoreRBAC: false,
		}
	)

	console.log('starting postgraphile')
	const app = require('express')()
	app.use(require('body-parser').json())
	app.use('/graphql/:queryHash', requestHandler)
	app.listen(5555)
}


const verifyToken = require('./auth.js')

async function requestHandler(req, res) {
	const method = req.method

	if (method === 'OPTIONS') {
		res.header('Access-Control-Max-Age', 86400)
		addCORSHeaders(res)
		return res.end()
	}

	const queryHash = req.params.queryHash

	let query
	let personId
	let variables
	let requiresAuth = false
	if (method === 'GET') {
		variables = req.query

		query = publicQueryMap[queryHash]
		if (!query) {
			requiresAuth = true
			query = secureQueryMap[queryHash]
			if (!query) return res.end(404)
		}
	}
	else if (method === 'POST') {
		variables = req.body.variables

		requiresAuth = true
		query = secureMutationMap[queryHash]
		if (!query) return res.end(404)
	}
	else
		return res.end(405)

	if (requiresAuth) {
		const [verifiedPersonId, statusCode] = verifyToken(req)
		if (statusCode) return res.end(statusCode)

		personId = verifiedPersonId
	}


	const role = personId
		? 'anonymous_user'
		: 'logged_in_user'

	const personIdFragment = personId
		? `, set_config('jwt.claims.person_id', ${personId}, true)`
		: ''

	const contextQuery = 'begin; '
		+ `select set_config('role', '${role}', true)`
		+ personIdFragment
		+ ';'

	const pgClient = await pgPool.connect()
	try {
		await pgClient.query(contextQuery)

		const { data, errors } = await executeGraphql(
		  graphqlSchema,
		  query, // fetched from the query map
		  null,
		  { pgClient },
		  variables,
		  null,
		)

		console.log(data)
		console.log(errors)

		// do something to end
		return res.json(data)
	}
	catch (error) {
		console.error(error)
		res.end(500)
	}
	finally {
		await pgClient.query('commit')
		await pgClient.release()
	}
}


function addCORSHeaders(res) {
  // res.setHeader('Access-Control-Allow-Origin', '*')
  res.setHeader('Access-Control-Allow-Origin', 'http://localhost:8080')
  // res.setHeader('Access-Control-Allow-Methods', 'HEAD, GET, POST')
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST')
  res.setHeader(
    'Access-Control-Allow-Headers',
    [
      'Origin',
      'X-Requested-With',
      // Used by `express-graphql` to determine whether to expose the GraphiQL
      // interface (`text/html`) or not.
      'Accept',
      // Used by PostGraphile for auth purposes.
      'Authorization',
      // Used by GraphQL Playground and other Apollo-enabled servers
      'X-Apollo-Tracing',
      // The `Content-*` headers are used when making requests with a body,
      // like in a POST request.
      'Content-Type',
      'Content-Length',
    ].join(', '),
  )
}


main()
