package service

import (
	"commercetools-ms-product/config"
	"context"
	"log"

	"github.com/labd/commercetools-go-sdk/platform"
	"golang.org/x/oauth2/clientcredentials"
)

func Connector() (*platform.ByProjectKeyRequestBuilder, context.Context) {

	client, err := platform.NewClient(&platform.ClientConfig{
		URL: config.Getenv("CT_API_URL"),
		Credentials: &clientcredentials.Config{
			TokenURL:     config.Getenv("CT_AUTH_URL") + "/oauth/token",
			ClientID:     config.Getenv("CT_CLIENT_ID"),
			ClientSecret: config.Getenv("CT_CLIENT_SECRET"),
			Scopes:       []string{config.Getenv("CT_SCOPE")},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	projectClient := client.WithProjectKey(config.Getenv("CT_PROJECT_KEY"))
	ctx := context.Background()

	return projectClient, ctx

}

func TestConnection() {
	projectClient, ctx := Connector()
	_, err := projectClient.Get().Execute(ctx)
	if err != nil {
		log.Fatalf("Error connecting to CommerceTools:\n%v", err)
	}
}

// // Get or Createa product type
// productTypeDraft := platform.ProductTypeDraft{
// 	Name: "a-product-type",
// 	Key:  ctutils.StringRef("a-product-type"),
// }

// productType, err := (projectClient.
// 	ProductTypes().
// 	WithKey(*productTypeDraft.Key).
// 	Get().
// 	Execute(ctx))

// if err != nil {
// 	if reqErr, ok := err.(platform.GenericRequestError); ok {
// 		if reqErr.StatusCode == 404 {
// 			productType, err = (projectClient.
// 				ProductTypes().
// 				Post(productTypeDraft).
// 				Execute(ctx))
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	} else {
// 		log.Fatal(err)
// 	}
// }

// r := rand.New(rand.NewSource(time.Now().UnixNano()))
// randomID := r.Int()
// productDraft := &platform.ProductDraft{
// 	Key: ctutils.StringRef(fmt.Sprintf("test-product-%d", randomID)),
// 	Name: platform.LocalizedString{
// 		"nl": "Een test product",
// 		"en": "A test product 2",
// 	},
// 	ProductType: platform.ProductTypeResourceIdentifier{
// 		ID: ctutils.StringRef(productType.ID),
// 	},
// 	Slug: platform.LocalizedString{
// 		"nl": fmt.Sprintf("een-test-product-%d", randomID),
// 		"en": fmt.Sprintf("a-test-product-%d", randomID),
// 	},
// }

// // The last argument is optional for reference expansion
// product, err := projectClient.
// 	Products().
// 	Post(*productDraft).
// 	WithQueryParams(
// 		platform.ByProjectKeyProductsRequestMethodPostInput{
// 			Expand: []string{"foobar"},
// 		},
// 	).
// 	Execute(ctx)

// if err != nil {
// 	log.Fatal(err)
// }

// // Alternatively you can pass query params via methods
// // product, err := projectClient.
// // 	Products().
// // 	// Get().Limit(1).
// // 	Post(*productDraft).
// // 	Expand([]string{"foobar"}).
// // 	Execute(ctx)

// fmt.Println(product)
// spew.Dump(product)

// 	productResults, err := projectClient.
// 		Products().
// 		Get().Limit(1).
// 		Execute(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	jsonBytes, err := json.Marshal(productResults)
// 	if err != nil {
// 		log.Fatalf("Failed to marshal product: %v", err)
// 	}

// 	fmt.Println(string(jsonBytes))

// }
