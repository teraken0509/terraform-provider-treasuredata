package treasuredata

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	td_client "github.com/treasure-data/td-client-go"
)

func TestAccTreasuredataSchedule_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTreasureDataScheduleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckTreasureDataScheduleConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"treasuredata_schedule.foobar", "name", "terraform_for_treasuredata_test_foobar"),
				),
			},
		},
	})
}

func testAccCheckTreasureDataScheduleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*td_client.TDClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "treasuredata_schedule" {
			continue
		}
		resp, _ := client.ListSchedules()
		for _, schedule := range *resp {
			if schedule.Name == rs.Primary.ID {
				return fmt.Errorf("Schedule still exists")
			}
		}
	}

	return nil
}

const testAccCheckTreasureDataScheduleConfigBasic = `
resource "treasuredata_schedule" "foobar" {
    name = "terraform_for_treasuredata_test_foobar"
}
`
