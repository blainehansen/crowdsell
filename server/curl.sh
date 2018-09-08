AUTH='Authorization: eyJpIjoiWk5XR292UG4iLCJlIjoxNTM2MzU1Mjk2fQ.dpX93Y0wEIZCtYvjFyhzL8-o4Q3NhMVxtX14Ge12rjk'
JSON='Content-Type: application/json'
SERVER='http://localhost:5050'


# curl -X POST $SERVER/create-user -H "$JSON" \
# 	-d '{"name": "dude", "email": "dude@gmail.com", "password": "pass"}'

# curl -X POST $SERVER/login \
# 	-d '{"email": "dude@gmail.com", "password": "pass"}'

# curl $SERVER/secure/user -H "$AUTH"


# curl -X POST $SERVER/secure/user/profile-image/sign -H "$AUTH"

curl -X "POST" "https://api.cloudinary.com/v1_1/crowdsell/image/upload" -H "$JSON" \
	-d "$(cat <<EOF
	{
		"file": "http://3.bp.blogspot.com/-S0KLwLVDy7o/TgDCP2GpfGI/AAAAAAAAFO8/1nPUnBZ9-_8/s1600/cool%2Bforest%2Bwallpapers%2B%252869%2529.jpg",
		"timestamp": 1536270375,
		"api_key": 856289479493379,
		"public_id": "ZNWGovPn",
		"signature": "d63d0ea99fe8a9ae8787cf7e78d72057fb58313f",
		"upload_preset": "profile"
	}
EOF
)"
