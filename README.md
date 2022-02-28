# GoChat!

GoChat! is a simple webapp that connects to a Redis server, and then uses WebSockets to send and receive messages.

## Features

- Chat history

## Quickstart

Install the dependencies:

```bash
make install
```

Then run the application:

```bash
make run
```

So you can see the chat app in action at: http://localhost:8080/chat

## Dependencies

- Golang 1.17+ (https://go.dev)
- Gin (https://gin-gonic.com)
- godotenv (https://github.com/joho/godotenv)
- Redis (https://redis.io)
- go-redis (https://redis.uptrace.dev)
- Gorilla WebSocket (https://github.com/gorilla/websocket)
- golangci-lint (https://golangci-lint.run)

## License

This project is licensed under the MIT license.

## Author

- Matheus Bosa
