FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/public ./public
ENV PORT=8080
EXPOSE 8080
CMD ["./server"]