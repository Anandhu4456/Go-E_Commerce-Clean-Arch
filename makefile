
.PHONY: run wire

run :
	docker compose up

wire :
	cd pkg/di && wire