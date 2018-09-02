AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM1ODQ0NjMwfQ.6bKFmgrGMAoicSj_ekUrfBwVeqaebMjT0QBhegj6iZc'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl $SERVER/projects

curl -X POST $SERVER/secure/user/card-token -H "$AUTH" -v

# curl -X "POST" "https://test.api.promisepay.com/token_auths" -H "$JSON" \
# 	-H "Authorization: Basic $ASSEMBLY_AUTH" \
# 	-d '{"token_type": "card", "user_id": "1"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl -X PATCH $SERVER/secure/user -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Guy", "bio": "Im a dude guy"}'

# curl -X POST $SERVER/secure/users/change-slug -H "$AUTH" -H "$JSON" \
# 	-d '{"urlSlug": "dude"}'

# curl -X POST $SERVER/secure/users/change-password -H "$AUTH" -H "$JSON" \
# 	-d '{"oldPassword": "pass", "newPassword": "dudepass"}'


# curl -X POST $SERVER/secure/profile-image/soemasdfdhash/png -H "$AUTH" \
# 	-F "file=@/home/blaine/Downloads/carbon.png"




# curl -X POST $SERVER/users/forgot-password -H "$JSON" \
# 	-d '{"email": "dude@gmail.com"}'

# curl -X POST $SERVER/users/recover-password -H "$JSON" \
# 	-d '{"recoveryToken": "ZHVkZUBnbWFpbC5jb206SFJyZWRaWG51dFdnVnRrMGNfZEhrUkNITXpNWEpxUkhIVHNO", "newPassword": "password"}'


# curl -X POST $SERVER/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff", "urlSlug": "dude-stuff"}'

# curl -X PATCH $SERVER/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'


# curl -X PATCH $SERVER/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"id": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'


# PROJECT_SLUG='vRAGNLVl'
# curl -X POST "$SERVER/secure/pledge/$PROJECT_SLUG" -H "$AUTH" -H "$JSON" -v \
# 	-d '{"Amount": 10000, "AccountId": "asdf", "IpAddress": "0.0.0.0", "DeviceId": "mac1234"}'
