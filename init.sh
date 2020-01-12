#!/bin/bash

go run /app/tls_status.go
crontab < /app/tls_cert.cron
httpd-foreground &
sleep infinity
