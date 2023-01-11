migrate-up:
	migrate -path migrations -database "mysql://root:root@/cryptoaml" -verbose up

migrate-down:
	migrate -path migrations -database "mysql://root:root@/cryptoaml" -verbose down
	
.PHONY: migrate-up migrate-down