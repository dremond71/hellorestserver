docker-compose up -d --always-recreate-deps --renew-anon-volumes --build
docker-compose ps
:: @timeout 5
:: docker logs hello-rest-server_mysqldb_1
