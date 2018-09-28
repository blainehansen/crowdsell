AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM4MTUzNTYyfQ.7LjppfzfpTZow6hr7PR5psZXKU112QlZ95gJVq0FCxo'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff"}'

PROJECT_SLUG="ZNWGovPn"
curl -X PATCH $SERVER/secure/projects/$PROJECT_SLUG -H "$AUTH" -H "$JSON" \
	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "promises": ["I will do this", "I will also do this"]}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl $SERVER/secure/user -H "$AUTH"



