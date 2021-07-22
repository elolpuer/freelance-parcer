FROM golang:1.15.8-alpine3.13

RUN mkdir -p /usr/src/app
ENV RabbitMQ=amqp://guest:guest@rabbitmq/
ENV Token=1643514464:AAFyDZX6c2pBy--T1Myp8rA-Lp_K0H7ZpGg
WORKDIR /usr/src/app

COPY . /usr/src/app

CMD ["go","run","main.go"]
