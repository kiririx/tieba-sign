FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./tieba-sign /tieba-sign

ENV bduss ""

USER 1000

WORKDIR /

ENTRYPOINT ["./tieba-sign"]