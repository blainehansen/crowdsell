AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTQwNzA5MzY0fQ.kWXdrzdfPEkL19-HLb1cV52Nodvz4StGlh5zH39ssz4'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'
GRAPHILE='http://localhost:5555'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

curl $GRAPHILE/graphql/FykmbGi6zUTHDLLffYX8H1IckNXUksfAT1edmghY5pg

# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff"}'


# curl -X PATCH $SERVER/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'

# PROJECT_SLUG="ZNWGovPnd"
# curl -X PATCH $SERVER/secure/projects/$PROJECT_SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"promises": [], "category": "stuff"}'


# curl $SERVER/secure/user -H "$AUTH"



