FROM golang:1.12.9-alpine3.10 as builder
COPY main.go .
RUN go build main.go
CMD ["./main"]