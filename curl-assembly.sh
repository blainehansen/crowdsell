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


# FEE_ID="70dcde6a-5a3d-4d46-99bc-40fb2844276f"


# curl -X "POST" "https://test.api.promisepay.com/items" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"id": "1",
# 		"name": "Seller's Project",
# 		"payment_type": "1",
# 		"seller_id": "1",
# 		"buyer_id": "2",
# 		"amount": 200000
# 	}
# EOF
# )"

# curl -X "PATCH" "https://test.api.promisepay.com/items/1/make_payment" -H "$AUTH" -H "$JSON" \
# 	-d "$(cat <<EOF
# 	{
# 		"account_id": "f3647c7f-1cb9-426d-a30a-5f2c70323525",
# 	}
# EOF
# )"
