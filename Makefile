# manually seed database creation
docker-up:
	docker compose up
docker-down:
	docker compose down --rmi all -v