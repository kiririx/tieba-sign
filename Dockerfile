FROM alpine:3

COPY ./tieba-sign /tieba-sign

ENV hour=6

ENV budss=""

USER 1000


CMD [ "./tieba-sign -h=$hour -b=$budss" ]