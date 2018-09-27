AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM3NDc0MTgxfQ.e0zA_dY9hmWthIkgNFdc6o_gPiT-ljtMV9AP4EzNeJU'
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

curl -X POST $SERVER/secure/projects/$PROJECT_SLUG/confirmation -H "$AUTH" -H "$JSON" \
	-d "$(cat <<EOF
	{
		"fulfills": {
			"proceed": true,
			"almostPromises": ["once", "other"],
			"commentary": "well stuff"
		}
	}
EOF
)"


# "unacceptable": {
# 	"fraudulentFlag": true,
# 	"brokenPromiseIds": [3, 4, 5]
# }
