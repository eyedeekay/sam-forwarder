FROM alpine:edge
ARG user=samcatd
ENV samhost=sam-host
ENV samport=7656
ENV args=""
ENV user=$user
ENV GOPATH=/usr
RUN apk update -U
RUN apk add go git make musl-dev webkit2gtk-dev gtk+3.0-dev
RUN mkdir -p /opt/$user
RUN adduser -h /opt/$user -D -g "$user,,,," $user
COPY . /usr/src/sam-forwarder
WORKDIR /usr/src/github.com/eyedeekay/sam-forwarder
RUN go get -u github.com/eyedeekay/sam-forwarder/samcatd
RUN make dylink install
USER $user
WORKDIR /opt/$user/
CMD samcatd -f /usr/src/eephttpd/etc/samcatd/tunnels.ini -littleboss start
