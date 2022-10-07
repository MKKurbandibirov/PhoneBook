all: up

up:
	docker compose -f ./docker-compose.yml up -d

ps:
	docker compose -f ./docker-compose.yml ps

down:
	docker compose -f ./docker-compose.yml down

fclean: down
	sudo rm -rf postgres