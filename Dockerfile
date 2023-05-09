FROM alpine:3

COPY ./tieba-sign /tieba-sign

USER 1000

CMD [ "./tieba-sign" ]