# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/FINTProsjektet/fint-tech-docs-service

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/FINTProsjektet/fint-tech-docs-service

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/fint-tech-docs-service

# Document that the service listens on port 8080.
EXPOSE 8080


#docker run -v $(pwd):/src --rm golang go get github.com/FINTProsjektet/fint-tech-docs-service/...