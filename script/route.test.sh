# signup
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "ljg"}' http://localhost:3000/api/v1/signup
curl -H "Content-Type:application/json" -X POST --data '{"username": "gjl", "password": "gjl"}' http://localhost:3000/api/v1/signup

# login
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "ljg"}' http://localhost:3000/api/v1/login
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "gjl"}' http://localhost:3000/api/v1/login
curl -H "Content-Type:application/json" -X POST --data '{"username": "lgj", "password": "gjl"}' http://localhost:3000/api/v1/login

# logout
# Auth wrapped, need session now
curl -H "Content-Type:application/json" -X POST  http://localhost:3000/api/v1/logout --cookie "SessionId=balabala"

# GET user
curl -X GET http://localhost:3000/api/v1/user?query="lj"
curl -X GET http://localhost:3000/api/v1/user?query=""

# GET pubKey
curl -X GET http://localhost:3000/api/v1/pubKey?username="ljg"

# POST pubKey
curl -H "Content-Type:application/json" -X POST  http://localhost:3000/api/v1/pubKey --cookie "SessionId=balabala"

# PUT pubKey
curl -H "Content-Type:application/json" -X PUT  http://localhost:3000/api/v1/pubKey --cookie "SessionId=balabala"

# DELETE pubKey
curl -H "Content-Type:application/json" -X PUT  http://localhost:3000/api/v1/pubKey --cookie "SessionId=balabala"
