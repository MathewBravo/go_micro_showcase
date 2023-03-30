FROM alpine:latest

RUN mkdir /app

COPY ./../gateway-service/gatewayApp /app

CMD ["/app/gatewayApp"]
