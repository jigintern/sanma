CONTAINER_NAME:=sanma-front

build:
	@docker build -t $(CONTAINER_NAME) .

run:
	@docker run --rm -it -p 8080:8080 -v ${PWD}:/app/ sanma-front

exec:
	@docker exec -it sanma-front sh

stop:
	@docker rm -f sanma-front
