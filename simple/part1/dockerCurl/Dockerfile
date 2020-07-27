FROM ubuntu:18.04

RUN apt update \
 && apt install -y curl dnsutils \
 && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /cwh

COPY ./net /cwh/

WORKDIR /cwh

CMD ["./net"]
