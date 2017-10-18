FROM golang:alpine AS go
WORKDIR /go/src/github.com/FINTprosjektet/fint-tech-docs-service
COPY go .
ENV CGO_ENABLED=0
ENV GOOS=linux
#RUN go install -a -v github.com/FINTprosjektet/fint-tech-docs-service
RUN go-wrapper download
RUN go-wrapper install

FROM node:alpine AS node
WORKDIR /src/client
COPY client .
RUN cd /src/client && yarn install && yarn buildClient
#RUN cd /src/client && npm install && npm run buildClient

FROM gradle:alpine
COPY config.yml config.yml
COPY --from=node /src/public public
COPY --from=go /go/bin/fint-tech-docs-service /usr/local/bin/ftds
USER root
RUN apk add --update tzdata && rm -rf /var/cache/apk/*
RUN chown -R gradle:gradle config.yml public
USER gradle
CMD ["/usr/local/bin/ftds"]
