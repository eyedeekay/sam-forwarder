FROM alpine:3.8
ARG user=samcatd
ARG path=example/path
ENV samhost=sam-host
ENV samport=7656
ENV args=""
ENV user=$user
RUN apk update -U
RUN apk add go git make musl-dev
RUN mkdir -p /opt/$user
RUN adduser -h /opt/$user -D -g "$user,,,," $user
COPY . /usr/src/samcatd
WORKDIR /usr/src/samcatd
RUN make deps full-test samcatd-web
RUN install -m755 bin/samcatd-web /usr/bin/samcatd-web
USER $user
WORKDIR /opt/$user/
CMD samcatd-web -f /usr/src/eephttpd/etc/eephttpd/eephttpd.conf #\
    #-s /opt/$user/ -sh=$samhost -sp=$samport $args
