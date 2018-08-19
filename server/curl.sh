AUTH='Authorization: eyJpIjoid3pYUWV3RG4iLCJlIjoxNTM0MDE5Mzg1fQ.2L8O6jA2KMKBN6G_kMIycXb2-UQEeC3OIoZKkdZVZfM'
JSON='Content-Type: application/json'
SLUG='wzXQewDn'

# curl -X POST http://localhost:5050/create-user \
# 	-d '{"name":"dude", "email":"dude@gmail.com", "password":"pass"}'
# curl -X POST http://localhost:5050/login \
# 	-d '{"email":"dude@gmail.com", "password":"pass"}'

# curl -X PATCH http://localhost:5050/secure/user -H "$AUTH" -H "$JSON" \
# 	-d '{"name":"Dude Guy", "bio":"Im a dude guy"}'

# curl -X POST http://localhost:5050/secure/users/change-slug -H "$AUTH" -H "$JSON" \
# 	-d '{"urlSlug":"dude"}'

# curl -X POST http://localhost:5050/secure/users/change-password -H "$AUTH" -H "$JSON" \
# 	-d '{"oldPassword": "pass", "newPassword": "dudepass"}'


# curl -X POST http://localhost:5050/secure/profile-image/soemasdfdhash/png -H "$AUTH" \
# 	-F "file=@/home/blaine/Downloads/carbon.png"




# curl -X POST http://localhost:5050/users/forgot-password -H "$JSON" \
# 	-d '{"email":"dude@gmail.com"}'

# curl -X POST http://localhost:5050/users/recover-password -H "$JSON" \
# 	-d '{"recoveryToken":"ZHVkZUBnbWFpbC5jb206SFJyZWRaWG51dFdnVnRrMGNfZEhrUkNITXpNWEpxUkhIVHNO", "newPassword":"password"}'


# curl -X POST http://localhost:5050/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name":"Dude Stuff", "description":"Various Dude Stuff", "urlSlug":"dude-stuff"}'

# curl -X PATCH http://localhost:5050/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name":"Changed Dude Stuff", "description":"Dude Stuff", "urlSlug":"dude-stuff"}'


# curl -X PATCH http://localhost:5050/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"id":"Changed Dude Stuff", "description":"Dude Stuff", "urlSlug":"dude-stuff"}'
