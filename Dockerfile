FROM alpine:3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a
ARG TARGETPLATFORM
RUN apk upgrade --no-cache
EXPOSE 9222
ENTRYPOINT ["/usr/bin/domain_exporter"]
COPY $TARGETPLATFORM/domain_exporter_*.apk /tmp/
RUN apk add --allow-untrusted /tmp/domain_exporter_*.apk
