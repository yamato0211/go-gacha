FROM golang:1.20-buster AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o api cmd/api/main.go

FROM gcr.io/distroless/base-debian11:latest

WORKDIR /

COPY --from=builder /app/api api
COPY --from=builder /app/.env .env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./api"]