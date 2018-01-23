package treasuredata

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	td_client "github.com/treasure-data/td-client-go"
)

func TestAccTreasuredataDatabase_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTreasureDataDatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckTreasureDataDatabaseConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"treasuredata_database.foobar", "name", "terraform_for_treasuredata_test_foobar"),
				),
			},
		},
	})
}

func testAccCheckTreasureDataDatabaseDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*td_client.TDClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "treasuredata_database" {
			continue
		}
		resp, _ := client.ListDatabases()
		for _, database := range *resp {
			if database.Name == rs.Primary.ID {
				return fmt.Errorf("Database still exists")
			}
		}

	}

	return nil
}

const testAccCheckTreasureDataDatabaseConfigBasic = `
resource "treasuredata_database" "foobar" {
    name = "terraform_for_treasuredata_test_foobar"
}
`
