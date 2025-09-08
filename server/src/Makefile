COMPOSE = docker compose

.PHONY: up down down-v start stop logs ps exec bash sh build restart rebuild

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

down-v:
	$(COMPOSE) down -v   ## ⚠️ Remove também volumes (apaga os dados!)

start:
	$(COMPOSE) start

stop:
	$(COMPOSE) stop

logs:
	$(COMPOSE) logs -f mariadb

ps:
	$(COMPOSE) ps

exec:
	$(COMPOSE) exec mariadb mysql -uroot -prootpass sys

bash:
	$(COMPOSE) exec mariadb bash || $(COMPOSE) exec mariadb sh

rebuild:
	$(COMPOSE) build --no-cache

restart:
	$(MAKE) stop
	$(MAKE) up
