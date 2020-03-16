d/up:
	@docker-compose up -d --build

d/down:
	@docker-compose down

# org
c-prune:
	docker container prune

i-prune:
	docker image prune

v-prune:
	docker volume prune

prune:
	make c-prune && make i-prune && make v-prune
