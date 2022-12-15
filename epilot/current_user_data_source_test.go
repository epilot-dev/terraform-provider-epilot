package epilot

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestCurrentUserDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "epilot_current_user" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.epilot_current_user.test", "email", "n.goel@epilot.cloud"),

					resource.TestCheckResourceAttr("data.epilot_current_user.test", "id", "placeholder"),
				),
			},
		},
	})
}
