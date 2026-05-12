FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gopher-slayer .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/gopher-slayer .
COPY --from=builder /app/frontend ./frontend
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/images ./images
EXPOSE 8080
CMD ["./gopher-slayer"]
