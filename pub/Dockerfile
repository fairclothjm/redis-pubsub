FROM golang:1.15

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod init pub
RUN go build -o pub

CMD ["/app/pub"]
