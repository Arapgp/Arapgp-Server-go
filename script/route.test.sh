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
curl -H "Content-Type:application/json" -X POST --data '{"pubKey": "lalala"}' http://localhost:3000/api/v1/pubKey

# PUT pubKey
curl -H "Content-Type:application/json" -X PUT  http://localhost:3000/api/v1/pubKey --cookie "SessionId=balabala"

# DELETE pubKey
curl -H "Content-Type:application/json" -X PUT  http://localhost:3000/api/v1/pubKey --cookie "SessionId=balabala"

# POST /user/:username/file
curl -H "Content-Type:application/json" -X POST --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/ljg/file
curl -H "Content-Type:application/json" -X POST --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/gjl/file

# GET /user/:username/file
curl -H "Content-Type:application/json" -X GET http://localhost:3000/api/v1/user/ljg/file?name="hahah.txt"

# PUT /user/:username/file
curl -H "Content-Type:application/json" -X PUT --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/ljg/file

# DELETE /user/:username/file
curl -H "Content-Type:application/json" -X DELETE --data '{"name": "hahah.txt"}' http://localhost:3000/api/v1/user/ljg/file
