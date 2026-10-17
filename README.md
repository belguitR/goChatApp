# goChatApp

A small real-time chat application with a Go WebSocket server and an Angular client. It demonstrates concurrent connections, message broadcasting, and a live online-user count.

## What works

- Multiple clients can join with a display name and exchange messages over `/ws`.
- The server exposes `/online` for the current connection count.
- The Go server serves the built Angular client from `public/browser` on port 8080 by default.

Messages and user names are kept in memory; there is no account system or persistent chat history. This is a learning project, not a production messaging service.

## Run locally

Requires Go 1.25 or newer. From the repository root:

```bash
go mod download
go run .
```

Open `http://localhost:8080` in two browser tabs to try a conversation. Set `PORT` to use a different port.

To rebuild the Angular client, use Node.js and run:

```bash
cd chatFront
npm ci
npm run build
```

The repository includes a prebuilt client in `public/browser` for the Go server. If you rebuild the client, copy its browser output there before running the server.

## Technical notes

The Go `Hub` tracks connected clients and broadcasts received messages. WebSocket handling uses `gorilla/websocket`; the Angular client connects to the server at the current page origin. The repository also includes a multi-stage Dockerfile for a compiled Go binary and the prebuilt client.
