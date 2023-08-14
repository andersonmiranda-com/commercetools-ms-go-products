# commercetools-ms-go-products

A product service meant to be a proxy between your application and [Commerce Tools products](https://docs.commercetools.com/api/projects/products) endpoints, writen in Go.

## Technologies
  
-   [Go 1.21](https://nodejs.org/docs/latest-v18.x/api/)
-   [Commercetools Go SDK](https://github.com/labd/commercetools-go-sdk/tree/main): SDK for Commercetools on Go
-   [FiberGo](https://github.com/gofiber/fiber): http Server
-   [Docker](https://www.docker.com/): For container generation

## Environment variables

The following variables must be defined/overwritten so that the service can work properly

| VARIABLE                  | DESCRIPTION                   | DEFAULT                         |
| ------------------------- | ----------------------------- | ------------------------------- |
| CT_API_URL                | commercetools API URL         | [https://api.commercetools.co](https://api.commercetools.co)  |
| CT_AUTH_URL               | commercetools auth URL        | [https://auth.commercetools.co](https://api.commercetools.co)  |
| CT_PROJECT_KEY            | commercetools project key     | -                               |
| CT_CLIENT_ID              | commercetools client id       | -                               |
| CT_CLIENT_SECRET          | commercetools client secret   | -                               |
| CT_SCOPE                  | commercetools scope           | -                               |

## Running local

Add a local .env file with the environment variables

Load dependency modules

`go mod download`

Run

`go run main.go`

## Using docker

### Build
```shell
docker build -t commercetools-ms-product .
docker run -p 4444:4444 --env-file ./.env commercetools-ms-product
```
