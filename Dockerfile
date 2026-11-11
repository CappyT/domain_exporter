FROM alpine:3@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf
ARG TARGETPLATFORM
RUN apk upgrade --no-cache
EXPOSE 9222
ENTRYPOINT ["/usr/bin/domain_exporter"]
COPY $TARGETPLATFORM/domain_exporter_*.apk /tmp/
RUN apk add --allow-untrusted /tmp/domain_exporter_*.apk
