package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kterada0509/terraform-provider-treasuredata/treasuredata"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: treasuredata.Provider})
}
