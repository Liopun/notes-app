FROM golang:1.19.5-alpine

RUN apk --no-cache add ca-certificates postgresql

COPY  ./ ./

# make it executable
RUN chmod +x postgres-ready.sh

RUN go mod download
RUN go build -o ./.bin/app ./cmd/app/main.go

CMD ["./.bin/app"]