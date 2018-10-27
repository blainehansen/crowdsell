const fs = require('fs')
const crypto = require('crypto')
const base64 = require('base64url')

const fileBases = ['public-queries', 'secure-queries', 'secure-mutations']

for (const fileBase of fileBases) {
	const fileName = `./client/queries/${fileBase}.json`

	const persistedQueries = JSON.parse(fs.readFileSync(fileName))

	const queryMap = {}
	for (const key of Object.keys(persistedQueries)) {
		const hash = base64.encode(crypto.createHash('sha256').update(key, 'utf8').digest())
		queryMap[key] = hash
	}

	fs.writeFileSync(fileName, JSON.stringify(queryMap))
}


