local-run:
	go run cmd/main.go

up:
	docker compose up

mysql-cli:
	docker exec -it mysql mysql -u root -p -D rinhaBackend

run-application-tests:
	go test ./internal/application/usecase/... -coverprofile=application-tests.out

show-application-coverage:
	go tool cover -html="application-tests.out"
