FROM alpine:3.8
RUN apk update -U
RUN apk add go git make musl-dev
RUN mkdir -p /opt/eephttpd
RUN adduser -h /opt/eephttpd -D -g 'eephttpd,,,,' eephttpd
COPY . /usr/src/eephttpd
WORKDIR /usr/src/eephttpd
RUN make deps server
RUN install -m755 bin/eephttpd /usr/bin/eephttpd

USER eephttpd
WORKDIR /opt/eephttpd/
VOLUME /opt/eephttpd/www
CMD eephttpd -sh=sam-host -sp=7656 -r
