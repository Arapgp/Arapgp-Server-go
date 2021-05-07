# signup
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "ljg"}' http://localhost:3000/api/v1/signup
curl -H "Content-Type:application/json" -X POST --data '{"username": "gjl", "password": "gjl"}' http://localhost:3000/api/v1/signup

# login
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "ljg"}' http://localhost:3000/api/v1/login
curl -H "Content-Type:application/json" -X POST --data '{"username": "ljg", "password": "gjl"}' http://localhost:3000/api/v1/login
curl -H "Content-Type:application/json" -X POST --data '{"username": "lgj", "password": "gjl"}' http://localhost:3000/api/v1/login