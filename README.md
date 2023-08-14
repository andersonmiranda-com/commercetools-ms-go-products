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

## Routes

```shell
GET / 
GET /:id
POST /
PUT /:id
PATCH /publish/:id
PATCH /unpublish/:id
```

## Request samples

#### GET

Makes a search for products

```shell
curl --location --globoff 'http://localhost:4444/api/products?limit=10&sort=[key]&where=(masterData(current(name(en%3D%22Bag%20%E2%80%9DAlexis%E2%80%9D%20medium%20Michael%20Kors%22))))'
```

Query parameters

```
	where
	expand
	sort
	limit
	offset
	withTotal
```


#### GET /:id

Get a product by id

```shell
curl --location 'http://localhost:4444/api/products/04f08c09-71fd-438f-aa58-f156d1d7fd9e?version=1&expend=productType'
```
Query parameters

```
	expand
```


#### POST /

```shell
curl --location --globoff 'http://localhost:4444/api/products?limit=10&sort=[key]&where=(masterData(current(name(en%3D%22Bag%20%E2%80%9DAlexis%E2%80%9D%20medium%20Michael%20Kors%22))))'
```


#### GET

```shell
curl --location --globoff 'http://localhost:4444/api/products?limit=10&sort=[key]&where=(masterData(current(name(en%3D%22Bag%20%E2%80%9DAlexis%E2%80%9D%20medium%20Michael%20Kors%22))))'
```


