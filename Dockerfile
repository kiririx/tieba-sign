FROM ubuntu:22.04

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER root

WORKDIR /

ENTRYPOINT ["./tieba-sign"]