FROM golang:1.15

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod init github.com/fairclothjm/redis-pubsub
WORKDIR /app/cmd/sub
RUN go build -o sub

CMD ["/app/cmd/sub/sub"]
