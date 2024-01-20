.PHONY: run

wire:
	cd pkg/di && wire

run:
	docker compose up