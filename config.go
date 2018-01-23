package treasuredata

import tdClient "github.com/treasure-data/td-client-go"

//Config ...
type Config struct {
	APIKey string
}

//NewClient ...
func (c *Config) NewClient() (*tdClient.TDClient, error) {
	client, err := tdClient.NewTDClient(tdClient.Settings{
		ApiKey:    c.APIKey,
		UserAgent: "Terraform for Treasure Data",
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
