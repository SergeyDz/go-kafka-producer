FROM golang:1.11 as builder
WORKDIR /app
COPY . /app
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-xorm/xorm
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get golang.org/x/net/context
RUN go get golang.org/x/oauth2/google
RUN go get google.golang.org/api/compute/v1
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine
WORKDIR /app
RUN apk --no-cache add ca-certificates && update-ca-certificates
ENV KAFKA_BROKER="localhost:9092"
ENV TOPIC="test"
COPY . /app
COPY --from=builder /app/app .
RUN ls /app
ENTRYPOINT ["./app"]