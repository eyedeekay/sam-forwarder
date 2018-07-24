# ephemeral-eepSite-SAM
An experiment to run an eepSite with an intentionally temporary destination
using the SAM bridge. Think the netcat web server trick, but with SAM.

## building

Just:

        make deps build

and it will be in the folder ./bin/

## usage

        ./bin/ephsite -addr=host:port

So, to serve an eepSite version of a local service on port 8080 -

        ./bin/ephsite -addr=127.0.0.1:8080
