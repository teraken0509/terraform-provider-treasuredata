package treasuredata

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	td_client "github.com/treasure-data/td-client-go"
)

func resourceTreasuredataSchedule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTreasuredataScheduleRead,
		Create: resourceTreasuredataScheduleCreate,
		Update: resourceTreasuredataScheduleUpdate,
		Delete: resourceTreasuredataScheduleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"cron": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"retry_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTreasuredataScheduleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	log.Printf("[DEBUG] Reading treasuredata schedule: %q", d.Id())

	resp, err := client.ListSchedules()
	if err != nil {
		return err
	}
	if len(*resp) == 0 {
		d.SetId("")
		return nil
	}
	for _, schedule := range *resp {
		if schedule.Name == d.Id() {
			d.Set("name", schedule.Name)
			d.Set("cron", schedule.Cron)
			d.Set("type", schedule.Type)
			d.Set("query", schedule.Query)
			d.Set("timezone", schedule.Timezone)
			d.Set("delay", schedule.Delay)
			d.Set("database", schedule.Database)
			d.Set("user_name", schedule.UserName)
			d.Set("priority", schedule.Priority)
			d.Set("retry_limit", schedule.RetryLimit)
			d.Set("result", schedule.Result)
			d.Set("next_time", schedule.NextTime)
			d.Set("created_at", schedule.CreatedAt)
			return nil
		}
	}

	return nil
}

func resourceTreasuredataScheduleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	name := d.Get("name").(string)
	options := getTreasuredataScheduleInput(d)

	resp, err := client.CreateSchedule(name, options)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] treasuredata schedule %q created.", resp.ID)

	d.SetId(resp.ID)
	d.Set("name", resp.Name)
	d.Set("cron", resp.Cron)
	d.Set("type", resp.Type)
	d.Set("query", resp.Query)
	d.Set("timezone", resp.Timezone)
	d.Set("delay", resp.Delay)
	d.Set("database", resp.Database)
	d.Set("user_name", resp.UserName)
	d.Set("priority", resp.Priority)
	d.Set("retry_limit", resp.RetryLimit)
	d.Set("result", resp.Result)
	d.Set("start", resp.Start)
	d.Set("created_at", resp.CreatedAt)

	return nil
}

func resourceTreasuredataScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	name := d.Get("name").(string)

	_, err := client.DeleteSchedule(name)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] treasuredata schedule %q deleted.", d.Id())

	d.SetId("")

	return nil
}

func resourceTreasuredataScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*td_client.TDClient)

	name := d.Get("name").(string)
	options := getTreasuredataScheduleInput(d)

	resp, err := client.UpdateSchedule(name, options)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] treasuredata schedule %q updated.", d.Id())

	d.SetId(resp.Name)
	d.Set("name", resp.Name)
	d.Set("cron", resp.Cron)
	d.Set("type", resp.Type)
	d.Set("query", resp.Query)
	d.Set("timezone", resp.Timezone)
	d.Set("delay", resp.Delay)
	d.Set("database", resp.Database)
	d.Set("user_name", resp.UserName)
	d.Set("priority", resp.Priority)
	d.Set("retry_limit", resp.RetryLimit)
	d.Set("result", resp.Result)
	d.Set("start", resp.Start)
	d.Set("created_at", resp.CreatedAt)

	return nil
}

func getTreasuredataScheduleInput(d *schema.ResourceData) map[string]string {
	var options = make(map[string]string)

	if v, ok := d.GetOk("cron"); ok {
		options["cron"] = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		options["type"] = v.(string)
	}
	if v, ok := d.GetOk("query"); ok {
		options["query"] = v.(string)
	}
	if v, ok := d.GetOk("timezone"); ok {
		options["timezone"] = v.(string)
	}
	if v, ok := d.GetOk("delay"); ok {
		options["delay"] = v.(string)
	}
	if v, ok := d.GetOk("database"); ok {
		options["database"] = v.(string)
	}
	if v, ok := d.GetOk("user_name"); ok {
		options["user_name"] = v.(string)
	}
	if v, ok := d.GetOk("priority"); ok {
		options["priority"] = v.(string)
	}
	if v, ok := d.GetOk("retry_limit"); ok {
		options["retry_limit"] = v.(string)
	}
	if v, ok := d.GetOk("result"); ok {
		options["result"] = v.(string)
	}

	return options
}
