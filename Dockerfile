FROM golang:1.11 as builder
WORKDIR /go/src/github.com/SergeyDz/go-kafka-producer
COPY . .
RUN go get github.com/Shopify/sarama 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine
WORKDIR /app
RUN apk --no-cache add ca-certificates && update-ca-certificates
ENV KAFKA_BROKER="localhost:9092"
ENV TOPIC="test"
COPY --from=builder /go/src/github.com/SergeyDz/go-kafka-producer/app .
ENTRYPOINT ["./app"]