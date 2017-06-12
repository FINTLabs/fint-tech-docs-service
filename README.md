# FINT technical docs services
Documentation service for FINT. You can see the service [here](https://docs.felleskomponent.no)

[![Go Report Card](https://goreportcard.com/badge/github.com/FINTprosjektet/fint-tech-docs-service)](https://goreportcard.com/report/github.com/FINTprosjektet/fint-tech-docs-service)

# Add a new repo to the service

* Add a webhook in the settings section of the github project. 
    * Payload URL: https://docs.felleskomponent.no/webhook
    
    | Parameter | Description |
    |-----------|-------------|
    | javadocs | `true` / `false` if the project has javadocs |
    | bintray | `false` / `false` if the project is on bintray |
    | lang | `java` or `net` |
    | fint_core_model | `false` / `false` |
    
    * Secret: You need to set the secret