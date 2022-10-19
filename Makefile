upload-image:
	docker build --tag molel/backend-face:latest .
	docker push molel/backend-face:latest

up:
	docker-compose up backend-face-service

down:
	docker-compose down

migration-test:
	docker-compose up --abort-on-container-exit integration

