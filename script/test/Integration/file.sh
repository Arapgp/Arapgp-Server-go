# this script for file

# 1. register
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "ljg"}' http://localhost:3000/api/v1/signup

# 2. upload PubKey
curl -H "Content-Type:application/json" -X POST --data '{"pubKey": "lalala"}' http://localhost:3000/api/v1/pubKey

# 3. curd PGPFile
# 3.1 create PGPFile
curl -H "Content-Type:application/json" -X POST --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/ljg/file
# reverse test
curl -H "Content-Type:application/json" -X POST --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/gjl/file
# 3.2 update PGPFile
curl -H "Content-Type:application/json" -X PUT --data '{"name": "hahah.txt", "pubKey": "lalala", "content": "-----BEGIN PGP MESSAGE-----\n\nasdsadasda-----END PGP MESSAGE-----"}' http://localhost:3000/api/v1/user/ljg/file
# 3.3 get PGPFile
curl -H "Content-Type:application/json" -X GET http://localhost:3000/api/v1/user/ljg/file?name="hahah.txt"
# 3.4 delete PGPFile
curl -H "Content-Type:application/json" -X DELETE --data '{"name": "hahah.txt"}' http://localhost:3000/api/v1/user/ljg/file
