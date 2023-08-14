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
