FROM golang:1.17-alpine AS builder
RUN mkdir /app
COPY main.go /app
WORKDIR /app
RUN go mod init main &&\
    go mod tidy &&\
    go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ /app
EXPOSE 8080

CMD ["./main"]