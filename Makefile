.PHONY:

# This make command will help you to running database for local development
db-up:
	docker compose up -d

# This make command will help you to stop running database for local development
db-down:
	docker compose down
