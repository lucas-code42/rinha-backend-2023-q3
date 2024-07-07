local-run:
	go run cmd/main.go

up:
	docker compose up

mysql-cli:
	docker exec -it mysql mysql -u root -p -D rinhaBackend
