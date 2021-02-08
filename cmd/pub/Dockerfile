FROM golang:1.15

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod init github.com/fairclothjm/redis-pubsub
WORKDIR /app/cmd/pub
RUN go build -o pub

CMD ["/app/cmd/pub/pub"]
