const connectionString = "postgres://user:asdf@localhost:5432/dev_database"

async function main() {
	const { Client } = require('pg')
	const client = new Client({ connectionString })

	await client.connect()

	await client.query(`begin; select set_config('jwt.claims.person_id', '1', true);`)

	try {
		const res = await client.query(`INSERT INTO stuff (stuff_text) values ('hello'); commit;`)
	}
	catch (e) {
		console.log(e)
	}
	finally {
		await client.end()
	}
}

main()
