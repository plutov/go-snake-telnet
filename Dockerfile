FROM golang:1.23-alpine AS builder
RUN apk add build-base
WORKDIR /root
COPY go.mod .
# COPY go.sum .
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o snakecmd main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata bash
WORKDIR /root
COPY --from=builder /root/snakecmd ./snakecmd

EXPOSE 8080

CMD ["./snakecmd"]
