d/up:
	@docker-compose up -d --build

d/down:
	@docker-compose down

d/re:
	make d/down && make d/up

# db
mysql:
	@cd db && \
	make mysql && \
	cd ..

db/init:
	@cd db && \
	make migrate/init && \
	make migrate/up && \
	cd ..

# org
c-prune:
	docker container prune

i-prune:
	docker image prune

v-prune:
	docker volume prune

prune:
	make c-prune && make i-prune && make v-prune
