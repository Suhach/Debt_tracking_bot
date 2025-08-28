GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
UTAG=user
MTAG=message
CONFIG_OAPI=./openapi/.openapi
OAPI_SCHEMA=./openapi/openapi.yaml
OAPI_OUTPUT_U=./internal/ogenerated/$(UTAG)/api.gen.go
OAPI_OUTPUT_M=./internal/ogenerated/$(MTAG)/api.gen.go
 
ifneq (,$(wildcard ./.env))
    include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@echo "  ${YELLOW}test            ${RESET} Run tests only"
	@echo "  ${YELLOW}lint            ${RESET} Run linters via golangci-lint"
	@echo "  ${YELLOW}tidy            ${RESET} Run tidy for go module to remove unused dependencies"
	@echo "  ${YELLOW}gen        	  ${RESET} Generate openapi"

.PHONY: gen 
gen:
	for tag in user message; do \
	rm -rf ./internal/ogenerated/$$tag/; \
	mkdir -p ./internal/ogenerated/$$tag/; \
	done

	oapi-codegen -config $(CONFIG_OAPI) -include-tags $(UTAG) -package $(UTAG) $(OAPI_SCHEMA) > $(OAPI_OUTPUT_U)
	oapi-codegen -config $(CONFIG_OAPI) -include-tags $(MTAG) -package $(MTAG) $(OAPI_SCHEMA) > $(OAPI_OUTPUT_M)

.PHONY: migrate-new
migrate-new:
	migrate create -ext sql -dir $(MIG_DIR)	${NAME}

.PHONY:	migrate
migrate:
	migrate -path $(MIG_DIR) -database $(DB_PATH) up

.PHONY: migrate-down
migrate-down:
	migrate -path $(MIG_DIR) -database $(DB_PATH) down

.PHONY: wire
wire:
	wire ./internal/app	