package treasuredata

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	//TreasureDataAPIKeyParamName ...
	TreasureDataAPIKeyParamName = "TD_API_KEY"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(TreasureDataAPIKeyParamName, nil),
				Description: "your Treasure Data APIKey",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"treasuredata_database": resourceTreasuredataDatabase(),
			"treasuredata_schedule": resourceTreasuredataSchedule(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey: d.Get("api_key").(string),
	}

	return config.NewClient()
}
