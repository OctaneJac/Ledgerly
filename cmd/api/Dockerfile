FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /app/api /api

EXPOSE 8080
CMD ["/api"]
