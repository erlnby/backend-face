FROM golang:1.19.1-alpine as modules
WORKDIR /modules
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.19.1-alpine
COPY --from=modules /go/pkg /go/pkg
WORKDIR /app
COPY . .

RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS=linux

CMD ["go", "test", "-v", "./integration-test/..."]