# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM gradle:alpine
MAINTAINER FINTProsjektet (www.fintprosjektet.no)


ADD fint-tech-docs-service ftds
COPY config.yml config.yml
COPY public public

USER root
RUN chown -R gradle:gradle config.yml ftds public

USER gradle

CMD ["./ftds"]