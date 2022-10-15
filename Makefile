.PHONY:

# This make command will help you to running database for local development
db-up:
	docker compose up -d

# This make command will help you to stop running database for local development
db-down:
	docker compose down

migrate-create:
	@read -p  "Migration name (eg:create_users, alter_entities, ...): " NAME; \
	goose -dir migrations create $$NAME sql

migrate-up:
	goose -dir migrations postgres "user=dot-library password=At0hlZ@vMNKMM2EYTShTRfDhlao53AR host=localhost port=25431 dbname=dot_library sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "user=dot-library password=At0hlZ@vMNKMM2EYTShTRfDhlao53AR host=localhost port=25431 dbname=dot_library sslmode=disable" down
