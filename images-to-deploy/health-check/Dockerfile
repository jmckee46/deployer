FROM scratch
MAINTAINER John McKee <john@subledger.com>

COPY health-check /
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY tls/files/* /tls/files/

EXPOSE 80 443
EXPOSE 81 444

LABEL com.subledger.deployer.push="true"

ENTRYPOINT ["/health-check"]
