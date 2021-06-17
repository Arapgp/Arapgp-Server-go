sudo docker-compose down
sudo docker-compose up -d
sleep 1
sudo docker exec arapgpdb bash -c "mongo ljgtest /setup/mongodb-setup.js"
