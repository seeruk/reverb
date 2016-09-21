FROM alpine:3.4
MAINTAINER Elliot Wright <hello@elliotdwright.com>

ADD ./dist/reverb-linux-amd64 /bin/reverb

EXPOSE 8080

CMD ["reverb"]
