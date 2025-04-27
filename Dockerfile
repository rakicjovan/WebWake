# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o wakeonlan-app

# Runtime Stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/wakeonlan-app .

EXPOSE 8080

CMD ["./wakeonlan-app"]
