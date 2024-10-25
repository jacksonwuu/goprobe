# goprobe

A high availability prober.

- [goprobe](#goprobe)
  - [1. Introduction](#1-introduction)
  - [2. Quick Start](#2-quick-start)
  - [3. Features](#3-features)
  - [4. Configure](#4-configure)
  - [n-1. Roadmap](#n-1-roadmap)
  - [n. License](#n-license)

## 1. Introduction

Goprobe is an active probing tool designed for _high availability_. 

## 2. Quick Start

```shell
./goprobe
```

You can manipulate probe job through Restful API:

```
GET /task

GET /tasks

POST /task

DELETE /task/:id
```

## 3. Features


Probers:

-   TCP
-   UDP
-   HTTP
-   TLS
-   ICMP
-   gRPC
-   WebSocket
-   DNS
-   SSH
-   Client
    -   Redis
    -   MySQL
    -   PostgreSQL
    -   Zookeeper
    -   Clickhouse
    -   Kafka
    -   and more...

Notifiers:

-   Slack
-   Email
-   Telegram
-   Discord
-   Log
-   SMS
-

Event:

-   ....

Metrics export:

-   Promethues

## 4. Configure


## n-1. Roadmap

You can view the project roadmap [here](./ROADMAP.md).

## n. License
