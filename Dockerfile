# base image
FROM uhub.service.ucloud.cn/yanser/centos7:latest

COPY consumer.tar.gz /home

RUN \
    tar -zxvf /home/consumer.tar.gz -C /home && \
    mkdir -p /home/consumer/log && \
    chmod +x /home/consumer/* && \
    true

WORKDIR /home/consumer

CMD [ "/home/consumer/run.sh" ]