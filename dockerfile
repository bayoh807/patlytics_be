# FROM ubuntu
FROM golang:1.23.2-alpine

WORKDIR /app
COPY . .

ARG API_KEY
ENV API_KEY=${API_KEY}

EXPOSE 8080
CMD ["go", "run", "main.go"]
