FROM alpine:3

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER root

WORKDIR /

ENTRYPOINT ["./tieba-sign"]