
.PHONY: run stop wire

run :
	docker compose up
stop:
	docker compose down
wire :
	cd pkg/di && wire