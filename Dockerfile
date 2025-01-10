FROM ubuntu:22.04

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER 1000

WORKDIR /

ENTRYPOINT ["./tieba-sign"]