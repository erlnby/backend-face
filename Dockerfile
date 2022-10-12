FROM golang:1.19.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o ./app ./cmd/main.go

EXPOSE 80

RUN addgroup -S noroot && adduser -S noroot -G noroot
USER noroot:noroot

CMD [ "./app" ]