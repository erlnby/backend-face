upload image:
	docker build --tag molel/backend-face:latest .
	docker push molel/backend-face:latest

up:
	docker-compose --env-file .env.example up

down:
	docker-compose down
