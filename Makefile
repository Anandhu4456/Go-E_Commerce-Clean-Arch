.PHONY: run

wire:
	cd pkg/di && wire

run:
	sudo docker run --network host -p 8080:8080 ecommerce