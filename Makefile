APP=goback
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgresql://golang:abcd1234@localhost/webgo?sslmode=disable&search_path=tiketz"
MIGRATIONS_DIR=./migrations
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


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