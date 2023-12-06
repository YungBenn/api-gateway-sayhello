docker.up:
	docker-compose up -d

docker.down:
	docker-compose down

docker.clean:
	docker-compose kill && docker-compose rm -f
	docker rmi hello-service:v1
	docker rmi user-service:v1