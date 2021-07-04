db:
	docker compose up database

server:
	docker compose up server

build:
	docker compose build

run:
	docker compose up

down:
	docker compose down

gen:
	sqlboiler --wipe psql

build-db:
	docker compose build database

build-server:
	docker compose build server