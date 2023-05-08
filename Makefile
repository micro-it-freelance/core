up:
	docker compose up --build

down:
	docker compose down

test:
	docker compose \
	-f $(PWD)/docker-compose.yml \
	-f $(PWD)/docker-compose.test.yml \
	--env-file $(PWD)/.env \
	up $(service) --exit-code-from $(service) --build

	docker compose logs