FROM alpine:3

WORKDIR /

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER 1000

CMD "./tieba-sign"