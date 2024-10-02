ARG VERSION
FROM sammobach/go:latest as build
LABEL maintainer="Sam Mobach <hello@sammobach.com>"
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN mkdir -p /app/bin &&go build -tags dev -ldflags "-s -w -X 'github.com/nilpntr/gitdesk-forwarder/cmd.version=$VERSION'" -o bin/gitdesk-forwarder github.com/nilpntr/gitdesk-forwarder
CMD ["/app/bin/gitdesk-forwarder"]

FROM alpine:3.20
LABEL maintainer="Sam Mobach <hello@sammobach.com>"
RUN mkdir /app
COPY --from=build /app/bin /app
WORKDIR /app
CMD ["/app/gitdesk-forwarder"]