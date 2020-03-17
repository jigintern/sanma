# dev
dev:
	@docker run -it -v ${PWD}:/go-api -w /go-api -p 8070:8080 --name="sanma-api-dev" golang:1.14.0-alpine3.11 sh

# run
build:
	@docker build --force-rm -t sanma-api .

run:
	@docker run --name="sanma-api-c" sanma-api

rmi:
	@docker rmi sanma-api
