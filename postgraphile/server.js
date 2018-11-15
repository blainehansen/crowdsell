const { parse: parseGraphql, execute: executeGraphql } = require('graphql')

function parseQueryMap(queryMap) {
	const newQueryMap = {}

	for (const [hash, query] of Object.entries(queryMap)) {
		newQueryMap[hash] = parseGraphql(query)
	}

	return newQueryMap
}

const fs = require('fs')
const publicQueryMap = parseQueryMap(JSON.parse(fs.readFileSync('/public-queries.json')))
const secureQueryMap = parseQueryMap(JSON.parse(fs.readFileSync('/secure-queries.json')))
const secureMutationMap = parseQueryMap(JSON.parse(fs.readFileSync('/secure-mutations.json')))

console.log(publicQueryMap)
console.log(secureQueryMap)
console.log(secureMutationMap)

const environment = require('./environment.js')

function createConnectionString({ user, password, host, port, database }) {
	return `postgres://${user}:${password}@${host}:${port}/${database}`
}

const serverConnectionConfig = {
	host: environment['DOCKER_DATABASE_HOST'],
	port: environment['DATABASE_PORT'],
	database: environment['DATABASE_DB_NAME'],

	user: 'postgraphile_server_user',
	password: environment['POSTGRAPHILE_DATABASE_PASSWORD'],
}

const inspectConnectionConfig = {
	...serverConnectionConfig,

	user: 'postgraphile_inspect_user',
	password: environment['POSTGRAPHILE_INSPECT_DATABASE_PASSWORD'],
}


// Create a postgres pool for efficiency
const pg = require("pg")
const pgPool = new pg.Pool({ connectionString: createConnectionString(serverConnectionConfig) })

let graphqlSchema


async function main() {
	// creates postgraphile schema
	const { createPostGraphileSchema } = require('postgraphile-core')
	const PgSimplifyInflectorPlugin = require("@graphile-contrib/pg-simplify-inflector")


	graphqlSchema = await createPostGraphileSchema(
		createConnectionString(inspectConnectionConfig),
		["public"],
		{
			dynamicJson: true,
			appendPlugins: [PgSimplifyInflectorPlugin],
			ignoreRBAC: false,
		}
	)

	console.log('starting postgraphile')
	const app = require('express')()

	app.get(
		'/graphql/:queryHash',
		handleOptions,
		publicRequestHandler,
	)

	app.use(
		'/graphql/secure/:queryHash',
		handleOptions,
		require('body-parser').json(),
		secureRequestHandler,
	)

	app.listen(5555)
}


const verifyToken = require('./auth.js')

async function publicRequestHandler(req, res) {
	// can only handle gets
	if (req.method !== 'GET') return res.status(405).end()

	const queryHash = req.params.queryHash
	const query = publicQueryMap[queryHash]
	if (!query) {
		console.error('404: ', queryHash)
		return res.status(404).end()
	}

	const variables = req.query

	return await handleQuery(res, query, variables, false)
}


async function secureRequestHandler(req, res) {
	const queryHash = req.params.queryHash

	let query
	let variables
	let isPatch = false
	switch (req.method) {
		case 'GET':
			variables = req.query

			query = secureQueryMap[queryHash]
			if (!query) return res.status(404).end()
			break

		case 'PATCH':
			isPatch = true
		case 'POST':
			variables = req.body.variables
			query = secureMutationMap[queryHash]
			if (!query) return res.status(404).end()
			break

		default:
			return res.status(405).end()
	}

	const [personId, statusCode] = verifyToken(req)
	if (statusCode) return res.status(statusCode).end()

	return await handleQuery(res, query, variables, isPatch, personId)
}

async function handleQuery(res, query, variables, isPatch, personId = undefined) {
	const contextQuery = personId
		? `begin; select set_config('role', 'postgraphile_known_user', true), set_config('jwt.claims.person_id', ${personId}, true)`
		: null

	let data, errors
	const pgClient = await pgPool.connect()
	try {
		if (contextQuery) await pgClient.query(contextQuery)

		const { data: requestData, errors: requestErrors = [] } = await executeGraphql(
			graphqlSchema,
			query, // fetched from the query map
			null,
			{ pgClient },
			variables,
			null,
		)

		data = requestData
		errors = requestErrors
	}
	catch (error) {
		errors = [error]
	}
	finally {
		if (contextQuery) await pgClient.query('commit')
		await pgClient.release()
	}

	if (errors.length > 0) {
		console.error(errors)
		return res.status(500).end()
	}
	else {
		if (isPatch) res.status(204).end()
		else return res.json(data)
	}

}


function handleOptions(req, res, next) {
	// if (req.method !== 'OPTIONS') return next()

	res.header('Access-Control-Max-Age', 86400)

	res.setHeader('Access-Control-Allow-Origin', '*')
	res.setHeader('Access-Control-Allow-Methods', 'OPTIONS, HEAD, GET, POST, PATCH')
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
			// // Used by GraphQL Playground and other Apollo-enabled servers
			// 'X-Apollo-Tracing',
			// The `Content-*` headers are used when making requests with a body,
			// like in a POST request.
			'Content-Type',
			'Content-Length',
		].join(', '),
	)

	return next()
}


main()
