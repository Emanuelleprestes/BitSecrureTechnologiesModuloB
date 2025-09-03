COMPOSE = docker compose -f docker-compose.yml

.PHONY: up down start stop logs ps exec bash

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

down-v:
	$(COMPOSE) down -v   # remove também os volumes (apaga os dados do banco)

start:
	$(COMPOSE) start

stop:
	$(COMPOSE) stop

logs:
	$(COMPOSE) logs -f mariadb

ps:
	$(COMPOSE) ps

exec:
	$(COMPOSE) exec mariadb mariadb -u$${MYSQL_USER:-root} -p$${MYSQL_PASSWORD:-rootpass} $${MYSQL_DATABASE:-sys}

bash:
	$(COMPOSE) exec mariadb bash
