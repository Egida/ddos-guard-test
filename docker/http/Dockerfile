FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o main ./cmd/http

EXPOSE 8080

CMD [ "/app/main" ]