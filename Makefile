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

lint:
	golangci-lint run