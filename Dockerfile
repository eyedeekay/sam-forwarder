FROM alpine:edge
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
RUN make deps samcatd
RUN install -m755 bin/samcatd /usr/bin/samcatd
USER $user
WORKDIR /opt/$user/
CMD samcatd -f /usr/src/eephttpd/etc/samcatd/eephttpd.conf
