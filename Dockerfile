FROM golang:1.19.1-alpine3.16 AS modules
WORKDIR /modules
COPY *.mod *.sum ./
RUN go mod download

FROM golang:1.19.1-alpine3.16 AS builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /app
COPY ./ ./
RUN go build -o ./app ./cmd/main.go
CMD [ "./app" ]