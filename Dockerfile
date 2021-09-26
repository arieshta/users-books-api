FROM golang:1.17.1-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./config /app/config
COPY ./controllers /app/controllers
COPY ./lib/database /app/lib/database
COPY ./middlewares /app/lib/middlewares
COPY ./models   /app/models
COPY ./routes /app/routes
COPY .env .
COPY main.go .

RUN go build -o program

EXPOSE 8000

CMD [ "/program" ]
