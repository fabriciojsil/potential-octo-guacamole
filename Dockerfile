FROM golang:1.10 as builder

MAINTAINER fabriciojsil@gmail.com

ADD . $GOPATH/src/github.com/fabriciojsil/potential-octo-guacamole
WORKDIR $GOPATH/src/github.com/fabriciojsil/potential-octo-guacamole

RUN make build

FROM debian:jessie-slim

COPY --from=builder go/src/github.com/fabriciojsil/potential-octo-guacamole/counting-request-server /bin/counting-request-server

ENTRYPOINT ["/bin/counting-request-server"]
