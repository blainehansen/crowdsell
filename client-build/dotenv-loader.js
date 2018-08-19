const dotenv = require('dotenv')
const dotenvExpand = require('dotenv-expand')

module.exports = function(source) {
	const buf = Buffer.from(source)
	const config = dotenvExpand(dotenv.parse(buf))

	throw "stuff"
	return `module.exports = (function() { return ${JSON.stringify(config)}; })();`
}
