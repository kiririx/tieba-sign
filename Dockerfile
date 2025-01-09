FROM alpine:3

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER 1000

WORKDIR /

CMD "./tieba-sign"