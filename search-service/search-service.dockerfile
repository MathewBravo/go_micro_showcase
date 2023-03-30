FROM alpine:latest

RUN mkdir /app

COPY ./../search-service/searchApp /app

CMD ["/app/searchApp"]
