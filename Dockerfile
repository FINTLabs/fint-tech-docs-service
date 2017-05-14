# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM gradle:alpine
MAINTAINER FINTProsjektet (www.fintprosjektet.no)

USER gradle

ADD fint-tech-docs-service ftds
COPY config.yml config.yml
COPY public public

CMD ["./ftds"]