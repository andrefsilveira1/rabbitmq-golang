# Rabbit MQ with Golang

This repository must only introduce the firsts steps with rabbit mq with Golang.

Requirements:

 [ x ] - RabbitMQ
 
 [ x ] - Golang


## What is RabbitMQ ?

RabbitMQ it is a open-source message delivery system that implements AMQP (Advancced Message Queuing Protocol)

### Important concepts of RabbitMQ that you should learn:

- Producer
- Queue
- Consumer
- Exchange

## What is Golang ?

A programming language :nerd_face:


## Instaling RabbitMQ:

First of all, install RabbitMQ in you OS, if you do not have it. There is some ways to do it:

- Download RabbitMQ: https://www.rabbitmq.com/download.html
- Docker: `docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management`
- Use Docker Compose (read the .yaml file)


# How execute this project:

- Instal amqp lib: go get github.com/streadway/amqp
- Execute `docker compose up -d` (run rabbitMQ service)
- Start the receiver at the folder `main`. Run: `go run receiver.go`
- Run main.go and type a message: `go run main.go`

This service will sent only one message per execution. Feel free do adjust as you want.

# Important:

This example it's a little Chat Messenger demonstration. But, RabbitMQ is designed for reliability first, but not for real-time message communication.

