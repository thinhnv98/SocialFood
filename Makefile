db:
	docker compose up database

run:
	docker compose up server

down:
	docker compose down

gen:
	sqlboiler --wipe psql