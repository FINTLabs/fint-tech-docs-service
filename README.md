# FINT technical docs services
Documentation service for FINT. You can see the service [here](https://docs.felleskomponent.no)

[![Go Report Card](https://goreportcard.com/badge/github.com/FINTprosjektet/fint-tech-docs-service)](https://goreportcard.com/report/github.com/FINTprosjektet/fint-tech-docs-service)

# Add a new repo to the service

* Add a webhook in the settings section of the github project. 
    * Payload URL: https://docs.felleskomponent.no/webhook?parameter=value
    
    | Parameter | Value | Default | Description |
    |-----------|-------------|-----------|-------------|
    | javadocs | `true` or `false`  | `false` | |
    | bintray | `false` or `false`  | `false` | |
    | lang | `java` or `net` | | |
    | fint_core_model | `false` or `false` | `false` | |
    
    * Secret: You need to set the secret