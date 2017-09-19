package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mioi/terraform-provider-nsone/nsone"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nsone.Provider,
	})
}
