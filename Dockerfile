FROM httpd:2.4.41

RUN apt update && apt upgrade -y
RUN apt install -y less procps golang cron openssl

RUN mkdir /app

ADD tls_status.go          /app/tls_status.go
ADD init.sh                /app/init.sh
ADD logo.png               /usr/local/apache2/htdocs/logo.png
ADD httpd.conf             /usr/local/apache2/conf/httpd.conf
ADD get_certs_from_domain  /usr/local/bin/get_certs_from_domain
ADD tls_cert.cron          /app/tls_cert.cron

RUN mkdir -p /usr/sbin/.cache/

RUN chmod +x               /app/init.sh
RUN chmod +x               /usr/local/bin/get_certs_from_domain
RUN chown daemon.daemon -R /usr/local/apache2/
RUN chmod 777              /usr/sbin/.cache

USER daemon

ENTRYPOINT ["/app/init.sh"]
