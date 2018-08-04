# eyJpIjoid3pYUWV3RG4iLCJlIjoxNTMzNDExMTQzfQ.fEKTIukfWIwoZhb-dZXUnqRDrDQc0P4fUgBdfqzvoEA

# curl -X POST http://localhost:5050/create-user -d '{"name":"dude", "email":"dude@gmail.com", "password":"pass"}'
# curl -X POST http://localhost:5050/create-user -d '{"name":"man", "email":"man@gmail.com", "password":"pass"}'
# curl -X POST http://localhost:5050/login -d '{"email":"dude@gmail.com", "password":"pass"}'

# curl -X POST http://localhost:5050/secure/users/change-slug \
# 	-H 'Content-Type: application/json' \
# 	-H 'Authorization: eyJpIjoid3pYUWV3RG4iLCJlIjoxNTMzNDExMTQzfQ.fEKTIukfWIwoZhb-dZXUnqRDrDQc0P4fUgBdfqzvoEA' \
# 	-d '{"url_slug":"dude"}'



# curl -X POST http://localhost:5050/users/forgot-password \
# 	-H 'Content-Type: application/json' \
# 	-d '{"email":"dude@gmail.com"}'

# curl -X POST http://localhost:5050/users/recover-password \
# 	-H 'Content-Type: application/json' \
# 	-d '{"recovery_token":"ZHVkZUBnbWFpbC5jb206SFJyZWRaWG51dFdnVnRrMGNfZEhrUkNITXpNWEpxUkhIVHNO", "new_password":"password"}'
