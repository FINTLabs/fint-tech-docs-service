FROM node:alpine AS build1
WORKDIR /src/client
ADD client /src/client
RUN cd /src/client && yarn install && yarn buildClient
#RUN cd /src/client && npm install && npm run buildClient

FROM golang:alpine AS build2
ADD . /go/src/github.com/FINTprosjektet/fint-tech-docs-service
WORKDIR /go/src/github.com/FINTprosjektet/fint-tech-docs-service
ENV CGO_ENABLED=0
ENV GOOS=linux 
RUN go build .

FROM gradle:alpine
ADD config.yml config.yml
COPY --from=build1 /src/public public
COPY --from=build2 /go/src/github.com/FINTprosjektet/fint-tech-docs-service/fint-tech-docs-service ftds
USER root
RUN chown -R gradle:gradle config.yml ftds public
USER gradle
CMD ["./ftds"]
