AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM2Nzk1NTU1fQ.VcVUoy8FyWiXuW8qioWe4T0i0WNSX3uSVYXgwdJbV-o'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl $SERVER/secure/user -H "$AUTH"


PROJECT_SLUG="ZNWGovPn"
# curl -X POST $SERVER/secure/project/$PROJECT_SLUG/uploads/sign -H "$AUTH" \
# 	-d '["fasdfa9sd8f79asdf", "dkfjkfkaf8d89sfad"]'


# {"objectName":"ZNWGovPn/dkfjkfkaf8d89sfad","signature":"3234eec76a95be1a6ea71af63850c86fd6c1de9c","timestamp":1536707063}
# {"objectName":"ZNWGovPn/fasdfa9sd8f79asdf","signature":"e5bd60c742ece6b632cb7012c4400f546bca4766","timestamp":1536707063}
curl -X POST $SERVER/secure/project/$PROJECT_SLUG/uploads/confirm -H "$AUTH" \
	-d "$(cat <<EOF
	[{
		"signature": "3234eec76a95be1a6ea71af63850c86fd6c1de9c",
		"timestamp": 1536707063,
		"hash": "dkfjkfkaf8d89sfad",
		"version": "1536707080"
	}, {
		"signature": "e5bd60c742ece6b632cb7012c4400f546bca4766",
		"timestamp": 1536707063,
		"hash": "fasdfa9sd8f79asdf",
		"version": "1536707080"
	}]
EOF
)"
