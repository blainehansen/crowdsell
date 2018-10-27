const fs = require('fs')
const environment = require('dotenv').parse(fs.readFileSync('/.env'))

const crypto = require('crypto')
const base64 = require('base64url')
const Hashids = require('hashids')


const hashids = new Hashids(environment.HASHID_SALT, parseInt(environment.HASHID_MIN_LENGTH), environment.HASHID_ALPHABET)
function decodeHashid(id) {
	return hashids.decode(id)[0]
}

function unixTime() {
	return Math.floor(Date.now() / 1000)
}

function getHmacBuffer(proposedEncodedToken) {
	const hmac = crypto.createHmac('sha256', 'some-key')

	return new Promise((resolve, reject) => {
		hmac.on('readable', () => {
		  const data = hmac.read()
		  if (data) resolve(data)
		})

		hmac.write(proposedEncodedToken)
		hmac.end()
	})
}

module.exports = async function verifyToken(req) {
	const token = req.header('Authorization')
	console.log(token)

	const tokenSegments = token.split('.')
	if (tokenSegments.length !== 2) return [false, 400]
	const [proposedEncodedToken, proposedEncodedSignature] = tokenSegments

	const proposedSignature = base64.toBuffer(proposedEncodedSignature)

	const actualSignature = await getHmacBuffer(proposedEncodedToken)

	if (!crypto.timingSafeEqual(actualSignature, proposedSignature)) return [false, 403]

	// decode the json
	const decodedToken = JSON.parse(base64.decode(proposedEncodedToken))

	// check the expiration
	if (decodedToken.e <= unixTime()) return [false, 401]

	// decode the hashid
	return [true, decodeHashid(decodedToken.i)]
}
