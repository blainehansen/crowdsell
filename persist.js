const fs = require('fs')
const loaderFunction = require('./client/queries/gql-loader.js')

for (const fileBase of ['public-queries', 'secure-queries', 'secure-mutations']) {
	// go through the graphql source files
	const sourceFilename = `./client/queries/${fileBase}.gql`
	// get the source
	const source = fs.readFileSync(sourceFilename, 'utf8')
	// pass it to the loader with the file name
	const finalString = loaderFunction(source, true)

	const targetFilename = `./client/queries/${fileBase}.json`
	fs.writeFileSync(targetFilename, finalString)
}
