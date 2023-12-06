docker.network:
	docker network create kong-net

docker.up:
	docker-compose up -d

docker.down:
	docker-compose down

docker.clean:
	docker compose down && docker network rm kong-net 