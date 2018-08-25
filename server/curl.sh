AUTH='Authorization: eyJpIjoid3pYUWV3RG4iLCJlIjoxNTM0MDE5Mzg1fQ.2L8O6jA2KMKBN6G_kMIycXb2-UQEeC3OIoZKkdZVZfM'
JSON='Content-Type: application/json'
SLUG='wzXQewDn'

# curl -X POST http://localhost:5050/create-user \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'
# curl -X POST http://localhost:5050/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl -X PATCH http://localhost:5050/secure/user -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Guy", "bio": "Im a dude guy"}'

# curl -X POST http://localhost:5050/secure/users/change-slug -H "$AUTH" -H "$JSON" \
# 	-d '{"urlSlug": "dude"}'

# curl -X POST http://localhost:5050/secure/users/change-password -H "$AUTH" -H "$JSON" \
# 	-d '{"oldPassword": "pass", "newPassword": "dudepass"}'


# curl -X POST http://localhost:5050/secure/profile-image/soemasdfdhash/png -H "$AUTH" \
# 	-F "file=@/home/blaine/Downloads/carbon.png"




# curl -X POST http://localhost:5050/users/forgot-password -H "$JSON" \
# 	-d '{"email": "dude@gmail.com"}'

# curl -X POST http://localhost:5050/users/recover-password -H "$JSON" \
# 	-d '{"recoveryToken": "ZHVkZUBnbWFpbC5jb206SFJyZWRaWG51dFdnVnRrMGNfZEhrUkNITXpNWEpxUkhIVHNO", "newPassword": "password"}'


# curl -X POST http://localhost:5050/secure/projects -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Dude Stuff", "description": "Various Dude Stuff", "urlSlug": "dude-stuff"}'

# curl -X PATCH http://localhost:5050/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"name": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'


# curl -X PATCH http://localhost:5050/secure/projects/$SLUG -H "$AUTH" -H "$JSON" \
# 	-d '{"id": "Changed Dude Stuff", "description": "Dude Stuff", "urlSlug": "dude-stuff"}'



AUTH="Authorization: Basic $ASSEMBLY_AUTH"

# curl "https://test.api.promisepay.com/users" -H "$AUTH"

# curl -X "POST" "https://test.api.promisepay.com/users" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"id": "1",
# 		"email": "seller@gmail.com",
# 		"first_name": "Seller",
# 		"country": "USA"
# 	}
# EOF
# )"

# ID="23e264ec-4686-4763-a36b-d5534f25f37a"
# curl -X "POST" "https://test.api.promisepay.com/bank_accounts" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"user_id": "1",
# 		"bank_name": "The Iron Bank",
# 		"account_name": "The Iron Bank - Seller",
# 		"routing_number": "324079555",
# 		"account_number": "12341234",
# 		"account_type": "checking",
# 		"holder_type": "personal",
# 		"country": "USA"
# 	}
# EOF
# )"

# curl -X "PATCH" "https://test.api.promisepay.com/users/1/disbursement_account" -H "$AUTH" -H "$JSON" \
# 	-d '{"account_id": "23e264ec-4686-4763-a36b-d5534f25f37a"}'




# curl -X "POST" "https://test.api.promisepay.com/users" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"id": "2",
# 		"email": "buyer@gmail.com",
# 		"first_name": "Buyer",
# 		"country": "USA"
# 	}
# EOF
# )"

# ID="f3647c7f-1cb9-426d-a30a-5f2c70323525"
# curl -X "POST" "https://test.api.promisepay.com/card_accounts" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"full_name": "Buyer User",
# 		"user_id": "2",
# 		"number": "4111111111111111",
# 		"expiry_month": "10",
# 		"expiry_year": "2020",
# 		"cvv": "123"
# 	}
# EOF
# )"


FEE_ID="70dcde6a-5a3d-4d46-99bc-40fb2844276f"


# curl -X "POST" "https://test.api.promisepay.com/items" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"id": "1",
# 		"name": "Seller's Project",
# 		"payment_type": "1",
# 		"seller_id": "1",
# 		"buyer_id": "8af719e680abb459c6bd75408b130f7d",
# 		"amount": 200000
# 	}
# EOF
# )"

curl -X "PATCH" "https://test.api.promisepay.com/items/1/make_payment" -H "$AUTH" -H "$JSON" \
	-d "$(cat <<EOF
	{
		"account_id": "f3647c7f-1cb9-426d-a30a-5f2c70323525",
	}
EOF
)"
