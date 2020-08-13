# FROM tmornini/http-spec:v0.9.0
FROM jmckee/http-spec:1.1
MAINTAINER Tom Mornini <tom@subledger.com>

COPY / /home
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /home

RUN apk add py3-ruffus

RUN pip3 install awscli --upgrade

LABEL com.subledger.deployer.push="true"

CMD ["/home/http-specs"]
