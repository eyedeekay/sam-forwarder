FROM alpine:edge
ARG user=samcatd
ENV samhost=sam-host
ENV samport=7656
ENV args=""
ENV user=$user
ENV GOPATH=/opt/$user/go
RUN apk update -U
RUN apk add go git make musl-dev
RUN mkdir -p /opt/$user
RUN adduser -h /opt/$user -D -g "$user,,,," $user
COPY . /usr/src/samcatd
WORKDIR /usr/src/samcatd
RUN go get -u github.com/eyedeekay/sam-forwarder/samcatd
RUN make all install
USER $user
WORKDIR /opt/$user/
CMD samcatd -f /usr/src/eephttpd/etc/samcatd/eephttpd.conf
