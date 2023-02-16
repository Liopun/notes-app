.PHONY:
.SILENT:
.DEFAULT_GOAL := run

gdraft:
	git add .
	git commit -m "${msg}"

git:
	git add .
	git commit -m "${msg}"
	git push

build.unix:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

build.win:
	go mod download && set CGO_ENABLED=0; go env -w GOARCH=amd64 GOOS=linux && go build -o ./.bin/app ./cmd/app/main.go

run: build.unix
	docker-compose up --remove-orphans app -d

dev:
	docker-compose up --build --remove-orphans dev -d

test.coverage:
	go tool cover -func=cover.out | grep "total"

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

lint:
	golangci-lint run

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/app/main.go