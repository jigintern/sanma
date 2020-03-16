build:
	docker build -t front .

run:
	docker run --rm -it -p 8080:8080 -v ${PWD}:/app/ front

exec:
	docker exec -it front sh

stop:
	docker rm -f front