FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /fds-backend ./cmd/api

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /fds-backend /app/fds-backend
COPY .env.example /app/.env.example
RUN mkdir -p /app/uploads/design-submissions /app/uploads/presets
EXPOSE 8080
CMD ["/app/fds-backend"]
