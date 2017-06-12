# FINT technical docs services
Documentation service for FINT. You can see the service [here](https://docs.felleskomponent.no)

[![Go Report Card](https://goreportcard.com/badge/github.com/FINTprosjektet/fint-tech-docs-service)](https://goreportcard.com/report/github.com/FINTprosjektet/fint-tech-docs-service)

# Add a new repo to the service

* Add a webhook in the settings section of the github project. 
    * Payload URL: https://docs.felleskomponent.no/webhook?parameter=value
    * Secret: *The secret*
    
## Query parameters

| Parameter | Value | Default | Description |
|-----------|-------------|-----------|-------------|
| javadocs | `true` or `false`  | `false` | Indicates if the project has JavaDocs |
| bintray | `false` or `false`  | `false` | Indicates if the project is on Bintray |
| lang | `java` or `net` | n/a | Indicates the language of the project. |
| fint_core_model | `false` or `false` | `false` | Indicates if the project is a core FINT model |
