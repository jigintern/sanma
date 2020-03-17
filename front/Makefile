CONTAINER_NAME:=sanma-front

build:
	@docker build -t $(CONTAINER_NAME) .

run:
	@docker run --rm -it -p 8080:8080 -v ${PWD}:/app/ --name="$(CONTAINER_NAME)" $(CONTAINER_NAME)

exec:
	@docker exec -it $(CONTAINER_NAME) sh

stop:
	@docker rm -f $(CONTAINER_NAME)
