const crypto = require('crypto')
const base64 = require('base64url')
const Hashids = require('hashids')

const environment = require('./environment.js')

const hashids = new Hashids(environment['HASHID_SALT'], parseInt(environment['HASHID_MIN_LENGTH']), environment['HASHID_ALPHABET'])
function decodeHashid(id) {
	return hashids.decode(id)[0]
}

function unixTime() {
	return Math.floor(Date.now() / 1000)
}

const signingKey = environment['SIGNING_KEY']
function getHmacBuffer(proposedEncodedToken) {
	const hmac = crypto.createHmac('sha256', signingKey)

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
	if (tokenSegments.length !== 2) return [null, 400]
	const [proposedEncodedToken, proposedEncodedSignature] = tokenSegments

	let proposedSignature, actualSignature
	try {
		proposedSignature = base64.toBuffer(proposedEncodedSignature)

		actualSignature = await getHmacBuffer(proposedEncodedToken)
	}
	catch (e) {
		console.error(e)
		return [null, 400]
	}

	if (!crypto.timingSafeEqual(actualSignature, proposedSignature)) return [null, 403]

	// decode the json
	let decodedToken
	try {
		decodedToken = JSON.parse(base64.decode(proposedEncodedToken))
	}
	catch (e) {
		return [null, 500]
	}

	// check the expiration
	if (decodedToken.e <= unixTime()) return [null, 401]

	// decode the hashid
	let decodedTokenId
	try {
		decodedTokenId = decodeHashid(decodedToken.i)
	}
	catch (e) {
		return [null, 500]
	}

	return [decodedTokenId, null]
}
