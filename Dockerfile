FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY baodou /usr/local/bin/baodou

ENTRYPOINT ["baodou"]
