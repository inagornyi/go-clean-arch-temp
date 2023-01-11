migrate-up:
	migrate -path migrations -database "mysql://root:root@/clean-arch-temp" -verbose up

migrate-down:
	migrate -path migrations -database "mysql://root:root@/clean-arch-temp" -verbose down
	
.PHONY: migrate-up migrate-down