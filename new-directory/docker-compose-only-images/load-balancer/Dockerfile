FROM ubuntu:16.04
MAINTAINER John McKee <john@subledger.com>

RUN useradd -ms /bin/bash nginx

RUN apt-get update
RUN apt-get -y dist-upgrade
RUN apt-get -yq install net-tools nginx
RUN apt-get install -y nginx-extras

COPY nginx.conf /etc/nginx/nginx.conf
COPY load-balancer.crt /etc/nginx/ssl/load-balancer.crt
COPY load-balancer.key /etc/nginx/ssl/load-balancer.key

ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
