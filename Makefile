.PHONY: db-reset
db-reset:
	@echo "Resetting DB"
	@goose down-to 0
	@goose up

.PHONY: seed
seed:
	@echo "Seeding..."
	@go run ./cmd/seed/.
	
