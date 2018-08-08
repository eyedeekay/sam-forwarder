FROM alpine:3.8
RUN apk update -U
RUN apk add go git make musl-dev
RUN mkdir -p /opt/eephttpd
RUN adduser -h /opt/eephttpd -D -g 'eephttpd,,,,' eephttpd

USER eephttpd
RUN git clone https://github.com/eyedeekay/sam-forwarder /opt/eephttpd/src
WORKDIR /opt/eephttpd/src
RUN make deps server

USER root
RUN ls -lah bin ; false
RUN cp bin/eephttpd /usr/bin/eephttpd
USER eephttpd

VOLUME /opt/eephttpd/www
CMD eephttpd
