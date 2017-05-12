# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM gradle:alpine
MAINTAINER FINTProsjektet (www.fintprosjektet.no)

# Copy the local package files to the container's workspace.
ADD fint-tech-docs-service ftds
COPY config.yml config.yml
COPY public public
RUN ls -lag .

CMD ["./ftds"]