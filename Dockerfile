# This is the template for the Dockerfile for golang project

FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o userapp .

#generated executable binary which is compatible for linux/amd64

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app .
EXPOSE 3015

CMD ["./userapp"]





# docker build -t go-app .

# docker run -p 4000:3015 go-app