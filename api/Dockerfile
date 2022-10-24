FROM golang:1.19.2-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
WORKDIR /app
COPY .env /app
COPY go.mod /app
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]