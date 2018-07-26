# sam-forwarder
Forward a local port to i2p over the SAM API.

## building

Just:

        make deps build

and it will be in the folder ./bin/

## usage

        ./bin/ephsite -addr=host:port

So, to serve an eepSite version of a local service on port 8080 -

        ./bin/ephsite -addr=127.0.0.1:8080
