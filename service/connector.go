package service

import (
	"commercetools-ms-product/config"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/labd/commercetools-go-sdk/platform"
	"golang.org/x/oauth2/clientcredentials"
)

type ctService struct {
	Connection *platform.ByProjectKeyRequestBuilder
}

type Service interface {
	Get(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

func NewService() Service {
	return &ctService{Connection: NewConnection()}
}

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

func NewConnection() *platform.ByProjectKeyRequestBuilder {

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

	return projectClient

}

func TestConnection() {
	projectClient, ctx := Connector()
	_, err := projectClient.Get().Execute(ctx)
	if err != nil {
		log.Fatalf("Error connecting to CommerceTools:\n%v", err)
	}
}
