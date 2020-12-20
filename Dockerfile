FROM alpine:latest
RUN mkdir /opt/cubicasa
RUN mkdir /etc/cubicasa
COPY ./bin/cubicasa /opt/cubicasa/
COPY ./config.json /etc/cubicasa/
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
CMD ["./opt/cubicasa/cubicasa"]