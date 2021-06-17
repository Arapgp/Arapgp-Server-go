sudo docker-compose down
sudo docker-compose up -d
sleep 1
sudo docker exec mongo bash -c "mongo ljgtest /setup/mongodb-setup.js"