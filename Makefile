APP=coffee-shop
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgresql://golang:golang@localhost/coffee_db?sslmode=disable&search_path=public"
MIGRATIONS_DIR=./migrations

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
# migrate create -ext sql -dir migrations/ -seq users > create file migrat berdasarkan nama tabel


install:
	go get -u ./... && go mod tidy

build:
	CGO_ENABLED=0 GOOS=linux go build -o ${BUILD}

test:
	go test -cover -v ./...

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0