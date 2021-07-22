FROM golang:1.15.8-alpine3.13

RUN mkdir -p /usr/src/app
ENV RabbitMQ=amqp://guest:guest@rabbitmq/
ENV PORT=5000
ENV HOST=
WORKDIR /usr/src/app

COPY . /usr/src/app

EXPOSE 5000

CMD ["go","run","server.go"]
