# Status website dashboard for TLS certificates

This is a containerized solution using a golang script and an Apache server
to display a list of certificates.

In order to use this software, your system needs the following:

* docker
* docker-compose

To build the image:

```
docker-compose build
```

To setup the image:

```
docker-compose up -d
```
