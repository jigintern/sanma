DB_NAME:=sanma
CONTAINER_NAME:=sanma-db
DB_HOST:=127.0.0.1
DB_PORT:=3306
DATASOURCE:=mysql://root:password@tcp($(DB_HOST):$(DB_PORT))/
MIGRATIONPATH:=migrations

mysql:
	@docker exec -it $(CONTAINER_NAME) mysql -u root -h $(DB_HOST) -P $(DB_PORT) --protocol tcp -p $(DB_NAME)

migrate/init:
	@docker exec -it $(CONTAINER_NAME) mysql -u root -h $(DB_HOST) -P $(DB_PORT) --protocol tcp -e "create database \`$(DB_NAME)\`" -p$(DB_PASSWORD)

migrate/new:
	@docker exec -it $(CONTAINER_NAME) migrate create -dir migrations -ext sql -seq "$(NAME)"

migrate/up:
	@make migrate_up

migrate/down:
	@make migrate_down

migrate_up:
	@docker exec -it $(CONTAINER_NAME) migrate -path $(MIGRATIONPATH) -database "$(DATASOURCE)$(DB_NAME)" up

migrate_down:
	@docker exec -it $(CONTAINER_NAME) migrate -path migrations -database "$(DATASOURCE)$(DB_NAME)" down 1

migrate/drop:
	@docker exec -it $(CONTAINER_NAME) mysql -u root -h $(DB_HOST) -P $(DB_PORT) --protocol tcp -e "drop database \`$(DB_NAME)\`" -p$(DB_PASSWORD)
