package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nokia/eda/apps/terraform-provider-fabrics/internal/provider"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "github.com/nokia-eda/fabrics-v1",
	}

	err := providerserver.Serve(context.Background(), provider.New("v1"), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
