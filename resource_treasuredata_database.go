package treasuredata

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	td_client "github.com/treasure-data/td-client-go"
)

func resourceTreasuredataDatabase() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTreasuredataDatabaseRead,
		Create: resourceTreasuredataDatabaseCreate,
		Delete: resourceTreasuredataDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTreasuredataDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	resp, err := client.ListDatabases()
	if err != nil {
		return err
	}
	if len(*resp) == 0 {
		d.SetId("")
		return nil
	}

	for _, database := range *resp {
		if database.Name == d.Id() {
			d.Set("name", database.Name)
			d.Set("count", database.Count)
			d.Set("created_at", database.CreatedAt.Format(time.RFC3339))
			d.Set("updated_at", database.UpdatedAt.Format(time.RFC3339))
			d.Set("permission", database.Permission)

			return nil
		}
	}
	return nil
}

func resourceTreasuredataDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	name := d.Get("name").(string)

	err := client.CreateDatabase(name, nil)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] treasuredata database %q created.", name)
	d.SetId(name)

	return resourceTreasuredataDatabaseRead(d, meta)
}

func resourceTreasuredataDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	err := client.DeleteDatabase(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] treasuredata database %q deleted.", d.Id())
	d.SetId("")

	return nil
}
