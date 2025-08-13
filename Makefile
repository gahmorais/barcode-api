run:
	go run cmd/api/main.go
run_prod:
	go run cmd/api/main.go -release=true
start_database:
	docker compose -f "docker/docker-compose.yml" up -d
stop_database:
	docker compose -f "docker/docker-compose.yml" down -v