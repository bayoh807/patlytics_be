# FROM ubuntu
FROM golang:1.23.2-alpine

WORKDIR /app
COPY . .

EXPOSE 8080
CMD ["go", "run", "main.go"]