run:
	go run cmd/main.go

docker-up:
	docker compose up -d

mysql-cli:
	docker exec -it mysql mysql -u root -p -D rinhaBackend
