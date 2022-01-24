FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]


