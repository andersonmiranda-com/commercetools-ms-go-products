# commercetools-ms-go-products

A product service designed as a proxy between your application and [commercetools Products API](https://docs.commercetools.com/api/projects/products). This service is built using Go.

## 🛠 Technologies Used

- [Go 1.21](https://golang.org/doc/go1.21)
- [Commercetools Go SDK](https://github.com/labd/commercetools-go-sdk/tree/main) - SDK for Commercetools in Go.
- [FiberGo](https://github.com/gofiber/fiber) - An efficient http server.
- [Docker](https://www.docker.com/) - For creating and managing containers.

## ⚙️ Configuration: Environment Variables

Ensure the following environment variables are set for the service to function correctly:

| VARIABLE                  | DESCRIPTION                   | DEFAULT VALUE                                         |
| ------------------------- | ----------------------------- | ----------------------------------------------------- |
| CT_API_URL                | commercetools API URL         | [https://api.commercetools.co](https://api.commercetools.co) |
| CT_AUTH_URL               | commercetools Auth URL        | [https://auth.commercetools.co](https://auth.commercetools.co) |
| CT_PROJECT_KEY            | commercetools Project Key     | Not set                                               |
| CT_CLIENT_ID              | commercetools Client ID       | Not set                                               |
| CT_CLIENT_SECRET          | commercetools Client Secret   | Not set                                               |
| CT_SCOPE                  | commercetools Scope           | Not set                                               |

## 🚀 Getting Started

### Running Locally:

1. Create a `.env` file at the root of the project and populate it with the required environment variables.
2. Install dependencies:

   ```bash
   go mod download
   ``````

3. Run the service

`go run main.go`

### Using docker

```bash
docker build -t commercetools-ms-product .
docker run -p 4444:4444 --env-file ./.env commercetools-ms-product
```

## 🚦 API Routes

```shell
GET /
GET /:id
POST /
PUT /:id
PATCH /publish/:id
PATCH /unpublish/:id
DELETE /:id
```



## 📝 API Usage Examples

### Query products

`GET /`

For more detailed information, refer to the official [commercetools documentation](https://docs.commercetools.com/api/projects/products#query-products).

#### Example Request

```
curl --location 'http://localhost:4444/api/products?limit=10&expand=productType&where=(masterData(current(name(en%3D%22Bag%20%E2%80%9DAlexis%E2%80%9D%20medium%20Michael%20Kors%22))))&sort=id&sort=key'
```

Accepted query parameters:

- where
- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand
- sort 
- limit
- offset
- withTotal


```
where, priceCurrency, priceCountry, priceCustomerGroup, priceChannel, localeProjection, expand, sort, limit, offset, withTotal    
```


### Get Product by ID

`GET /:id`

Refer to [Commercetools documentation](https://docs.commercetools.com/api/projects/products#get-product-by-id) for more details.

#### Example Request

```
curl --location 'http://localhost:4444/api/products/29bae503-3058-400a-b0f6-db7aaf2aae11'
```

Accepted query parameters:

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand   


### Create Product

`POST /`

Refer to  [Commercetools documentation](https://docs.commercetools.com/api/projects/products#create-product) for more details.

#### Example Request

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

Accepted query parameters:

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand             


### Update Product by ID

`PUT /:id`

Refer to [Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id) for more details.

#### Example Request

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

Accepted query parameters

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand


### Publish a Product by ID

`PATCH /publish/:id`

Get the last product version and send an update with `"action" : "publish"`

Refer to [Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id) for more details.

#### Example Request

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

Accepted query parameters

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand


### Unpublish a Product by ID

`PATCH /unpublish/:id`

Get the last product version and send an update with `"action" : "unpublish"`

Refer to [Commercetools documentation](https://docs.commercetools.com/api/projects/products#update-product-by-id) for more details.

#### Example Request

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

Accepted query parameters

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand


### Delete product by ID

`DELETE /:id`

Refer to [Commercetools documentation](https://docs.commercetools.com/api/projects/products#delete-product-by-id) for more details.

#### Example Request

```
curl --location --request DELETE 'http://localhost:4444/api/products/ccaf601d-169e-493e-96b1-47741d37df8f'
```

Accepted query parameters

- priceCurrency     
- priceCountry      
- priceCustomerGroup
- priceChannel      
- localeProjection   
- expand
- version (if ommited, the last version will be obtained from product)


