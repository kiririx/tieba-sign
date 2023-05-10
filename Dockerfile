FROM alpine:3

WORKDIR /

COPY ./ ./

ENV hour=6

ENV budss=""

USER 1000

CMD [ "./tieba-sign -h $hour -b $budss" ]