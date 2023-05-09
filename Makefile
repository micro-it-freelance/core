up:
	docker compose up --build

down:
	docker compose down

test:
	docker compose -f ./docker-compose.yml -f ./docker-compose.test.yml \
	up $(service) --build --exit-code-from $(service)

logs:
	@bash $(PWD)/submodules/core/logs.sh
