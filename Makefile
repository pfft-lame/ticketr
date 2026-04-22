.PHONY: db-reset
db-reset:
	@echo "Resetting DB"
	@goose down-to 0
	@goose up

.PHONY: db-down
db-down:
	@echo "Applying all DOWN migrations\n"
	@goose down-to 0

.PHONY: seed
seed:
	@echo "Seeding..."
	@go run ./cmd/seed/.
	
