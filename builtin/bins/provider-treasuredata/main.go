package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/kterada0509/terraform-provider-treasuredata"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: treasuredata.Provider,
	})
}
