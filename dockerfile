# FROM ubuntu
FROM golang:1.23.2-alpine AS builder

COPY . .

EXPOSE 8000
CMD ["go run main.go"]
