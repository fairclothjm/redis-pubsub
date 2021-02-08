FROM golang:1.15

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod init sub
RUN go build -o sub

CMD ["/app/sub"]
