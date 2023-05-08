up:
	docker compose up --build

down:
	docker compose down

test:
	docker compose -f ./docker-compose.yml -f ./docker-compose.test.yml \
	up --exit-code-from $(service) --build
