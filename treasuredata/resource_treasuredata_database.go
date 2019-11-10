package treasuredata

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceTreasuredataDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	log.Printf("[DEBUG] Reading treasuredata database: %s", d.Id())
	if _, ok := d.GetOk("name"); !ok {
		d.Set("name", d.Id())
	}

	databaseElement, err := client.ShowDatabase(d.Id())
	if err != nil {
		if err == fmt.Errorf("Database '%s' does not exist", d.Id()) {
			log.Printf("[WARN] Database (%s) does not exist", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("count", databaseElement.Count)
	d.Set("created_at", databaseElement.CreatedAt.Format(time.RFC3339))
	d.Set("updated_at", databaseElement.UpdatedAt.Format(time.RFC3339))
	d.Set("permission", databaseElement.Permission)
	d.Set("delete_protected", databaseElement.DeleteProtected)

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
