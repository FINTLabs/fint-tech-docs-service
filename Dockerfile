# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER FINTProsjektet (www.fintprosjektet.no)

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/FINTProsjektet/fint-tech-docs-service

ENV GRADLE_VERSION 3.3

RUN apt-get update
RUN apt-get -y install openjdk-8-jdk wget unzip

# Setup certificates in openjdk-8
RUN /var/lib/dpkg/info/ca-certificates-java.postinst configure

# Set path
ENV PATH ${PATH}:/usr/local/gradle-$GRADLE_VERSION/bin:/usr/local/node-v$NODE_VERSION-linux-x64/bin

# Install gradle
RUN wget https://services.gradle.org/distributions/gradle-$GRADLE_VERSION-bin.zip && \
    unzip gradle-$GRADLE_VERSION-bin.zip && \
    rm -f gradle-$GRADLE_VERSION-bin.zip

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/FINTProsjektet/fint-tech-docs-service

COPY public /go/bin/public
COPY config.yml /go/bin/ 
# Run the outyet command by default when the container starts.
WORKDIR /go/bin

ENTRYPOINT fint-tech-docs-service









#docker run -v $(pwd):/go/bin --rm \
#  golang:onbuild go get github.com/FINTProsjektet/fint-tech-docs-service/...




