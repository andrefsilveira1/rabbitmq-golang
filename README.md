# Rabbit MQ with Golang

This repository must only introduce the firsts steps with rabbit mq with Golang


## What is RabbitMQ ?

RabbitMQ it is a open-source message delivery system that implements AMQP (Advancced Message Queuing Protocol)

### Important concepts of RabbitMQ that you should learn:

- Producer
- Queue
- Consumer
- Exchange

## Instaling RabbitMQ:

First of all, install RabbitMQ in you OS. There is some ways to do it:

- Download RabbitMQ: https://www.rabbitmq.com/download.html
- Docker: `docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management`

- Use Docker Compose (read the file)