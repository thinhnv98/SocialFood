db:
	docker compose up database

server:
	docker compose up server

run:
	docker compose up

down:
	docker compose down

gen:
	sqlboiler --wipe psql