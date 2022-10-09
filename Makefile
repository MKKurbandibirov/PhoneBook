all: up

app:
	go run cmd/api/main.go

up:
	docker compose -f ./docker-compose.yml up -d

ps:
	docker compose -f ./docker-compose.yml ps

down:
	docker compose -f ./docker-compose.yml down

fclean: down
	sudo rm -rf postgres