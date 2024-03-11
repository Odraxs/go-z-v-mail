FROM golang:1.22.1-alpine3.19 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY . .

RUN go build -o bin/main
EXPOSE 3001

CMD ["./bin/main"]
