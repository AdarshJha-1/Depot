FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o depot ./cmd/depot

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/depot .
COPY web ./web

EXPOSE 6969

CMD ["./depot"]