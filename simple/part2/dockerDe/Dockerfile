FROM ubuntu:18.04

RUN mkdir -p /receiver

WORKDIR /receiver

COPY ./decrypt ./

RUN echo "ver2!"

CMD ["./decrypt"]
