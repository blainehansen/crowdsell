AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM5Mzc0Njc1fQ.QfrTb6QH1wYAVbpp5PS5WtPk-G4VbGjDcM449JbM1AQ'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff"}'


# curl -X PATCH $SERVER/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'

# PROJECT_SLUG="ZNWGovPn"
# curl -X PATCH $SERVER/secure/projects/$PROJECT_SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"promises": [], "category": "stuff"}'


# curl $SERVER/secure/user -H "$AUTH"



