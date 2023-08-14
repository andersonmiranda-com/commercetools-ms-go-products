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

## Running

### Local

Add a local .env file with the environment variables

Load dependency modules

`go mod download`

Run

`go run main.go`

### Using docker

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
DELETE /:id
```

---

## Request Examples

### Query products

`GET /`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#query-products)

```
curl --location 'http://localhost:4444/api/products?limit=10&expand=productType&where=(masterData(current(name(en%3D%22Bag%20%E2%80%9DAlexis%E2%80%9D%20medium%20Michael%20Kors%22))))&sort=id&sort=key'
```

Query parameters

```
    where              
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
    sort               
    limit          
    offset         
    withTotal       
```


### Get Product by ID

`GET /:id`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#get-product-by-id)


```
curl --location 'http://localhost:4444/api/products/29bae503-3058-400a-b0f6-db7aaf2aae11'
```

Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
```


### Create Product

`POST /`


[Commercetools documentation](https://docs.commercetools.com/api/projects/products#create-product)

```
curl --location 'http://localhost:4444/api/products?expand=productType' \
--header 'Content-Type: application/json' \
--data '{
  "productType" : {
    "id" : "e59a4778-c448-42a6-98ba-504a8016b8e9",
    "typeId" : "product-type"
  },
  "categories" : [ {
    "typeId" : "category",
    "key" : "c4"
  } ],
  "name" : {
    "en" : "Some Product"
  },
  "slug" : {
    "en" : "product_slug_987987987"
  },
  "masterVariant" : {
    "sku" : "SKU-1",
    "prices" : [ {
      "value" : {
        "currencyCode" : "EUR",
        "centAmount" : 4200
      }
    } ],
    "images" : [ {
      "url" : "http://my.custom.cdn.net/master.png",
      "label" : "Master Image",
      "dimensions" : {
        "w" : 303,
        "h" : 197
      }
    } ]
  },
  "variants" : [ {
    "images" : [ {
      "url" : "http://my.custom.cdn.net/variant.png",
      "label" : "Variant Image",
      "dimensions" : {
        "w" : 303,
        "h" : 197
      }
    } ]
  } ]
}'
```

Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
```


### Update Product by ID

`PUT /:id`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id)

```
curl --location --request PUT 'http://localhost:4444/api/products/04427947-f75b-4fde-995f-3d0a0d27c541' \
--header 'Content-Type: application/json' \
--data '{
  "version" : 2,
  "actions" : [ {
    "action" : "setDescription",
    "description" : {
      "en" : "The best product ever!"
    }
  } ]
}'
```

Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
```

### Publish a Product by ID

`PATCH /publish/:id`

Get the last product version and send an update with `"action" : "publish"`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id)

```
curl --location --request PATCH 'http://localhost:4444/api/products/publish/04427947-f75b-4fde-995f-3d0a0d27c541'
```

Equivalent to:

```
curl --location --request PUT 'http://localhost:4444/api/products/04427947-f75b-4fde-995f-3d0a0d27c541' \
--header 'Content-Type: application/json' \
--data '{
  "version" : 2,
  "actions" : [ {
    "action" : "publish",
  } ]
}'
```


Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
```

### Unpublish a Product by ID

`PATCH /unpublish/:id`

Get the last product version and send an update with `"action" : "unpublish"`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id)

```
curl --location --request PATCH 'http://localhost:4444/api/products/unpublish/04427947-f75b-4fde-995f-3d0a0d27c541'
```

Equivalent to:

```
curl --location --request PUT 'http://localhost:4444/api/products/04427947-f75b-4fde-995f-3d0a0d27c541' \
--header 'Content-Type: application/json' \
--data '{
  "version" : 2,
  "actions" : [ {
    "action" : "unpublish",
  } ]
}'
```

Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand             
```

### Delete product by ID

`DELETE /:id`

[Commercetools documentation](https://docs.commercetools.com/api/projects/products#delete-product-by-id)

```
curl --location --request DELETE 'http://localhost:4444/api/products/ccaf601d-169e-493e-96b1-47741d37df8f'
```

Query parameters

```
    priceCurrency     
    priceCountry      
    priceCustomerGroup
    priceChannel      
    localeProjection   
    expand
    version (if ommited, the last version will be obtained from product)
```


