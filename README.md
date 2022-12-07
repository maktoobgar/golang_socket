# Public Chat App

This project is simply a public chat app written in golang.\
People can create rooms, delete rooms and join rooms.

## Setup

**Download Dependencies**:
```
go mod download
```

**Start Server**:
```
go run main.go
```

## Technical Stand Point

This project is a combination of [gorrila websocket] and [gin framework] all combined together.
Most base codes come from a [example chat] app code from [gorrila websocket] project which I just developed a little more and combined it with [gin framework].


[gorrila websocket]: https://github.com/gorilla/websocket
[example chat]: https://github.com/gorilla/websocket/tree/master/examples/chat
[gin framework]: https://github.com/gin-gonic/gin
