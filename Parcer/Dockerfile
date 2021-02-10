FROM golang:1.15.8-alpine3.13

RUN mkdir -p /usr/src/app
ENV RabbitMQ=
WORKDIR /usr/src/app

COPY . /usr/src/app

CMD ["go","run","main.go"]
