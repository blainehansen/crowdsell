AUTH='Authorization: eyJpIjoiZ214Ymtud2UiLCJlIjoxNTQyMjQxNTM4fQ.Oz0UFQnjsHYqhtSlEaQRX2J46zb9h2R7G1AGPwgMrp0'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'
GRAPHILE='http://localhost:5555'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

curl $GRAPHILE/graphql/Or27CPXH319u1hnCSF5ck0nTVMXBYfHztg4TbB2AG4Y

# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff"}'


# curl -X PATCH $SERVER/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'

# PROJECT_SLUG="ZNWGovPnd"
# curl -X PATCH $SERVER/secure/projects/$PROJECT_SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"promises": [], "category": "stuff"}'


# curl $SERVER/secure/user -H "$AUTH"


curl http://localhost:5050/new-email -d '{ "email": "dude@gmail.com" }'
curl http://localhost:5050/validate -d '{ "validationToken": "24ugPcHPzKSS5vQuYT18zF8xokWupSWxOJvkOd1i1833_gnPXQba5vi_fWYWP1t22c-yobxhH_SWe5d8ENNXmQ" }'
