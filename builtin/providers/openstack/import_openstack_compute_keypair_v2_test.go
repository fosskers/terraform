package openstack

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccOpenStackComputeV2Keypair_importBasic(t *testing.T) {
	resourceName := "openstack_compute_keypair_v2.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2KeypairDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2Keypair_basic,
			},

			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}
