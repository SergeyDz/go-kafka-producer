# build stage
FROM golang:onbuild AS build-env
ADD . /src
RUN cd /src && go get github.com/Shopify/sarama
RUN cd /src && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
ENV KAFKA_BROKER ""
ENV TOPIC ""
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp