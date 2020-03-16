# dev
dev:
	@docker run -it -v ${PWD}:/go-api -w /go-api golang:1.14.0-alpine3.11 sh

# run
build:
	@docker build --force-rm -t sanma-api .

run:
	@docker run --name="sanma-api-c" sanma-api

rmi:
	@docker rmi sanma-api

# org
c-prune:
	docker container prune

i-prune:
	docker image prune

v-prune:
	docker volume prune

prune:
	make c-prune && make i-prune && make v-prune
