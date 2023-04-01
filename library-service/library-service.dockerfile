FROM alpine:latest

RUN mkdir /app

COPY ./../library-service/libraryApp /app

CMD ["/app/libraryApp"]
