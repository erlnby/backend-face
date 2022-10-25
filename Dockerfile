FROM golang:1.19.1-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/app ./cmd/main.go

FROM scratch AS run

COPY --from=build /app/app /app
COPY --from=build /app/migrations /migrations

ENTRYPOINT [ "/app" ]