// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsNodeAttWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsNodeAttDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "target_dn", "topology/pod-1/node-101"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "encapsulation", "vlan-101"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_leaf.test", "mode", "regular"),
				),
			},
			{
				Config:      testConfigFvRsNodeAttNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_static_leaf data source"),
			},
		},
	})
}

const testConfigFvRsNodeAttDataSourceDependencyWithFvAEPg = testConfigFvRsNodeAttMinDependencyWithFvAEPg + `
data "aci_relation_to_static_leaf" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/node-101"
  depends_on = [aci_relation_to_static_leaf.test]
}
`

const testConfigFvRsNodeAttNotExistingFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
data "aci_relation_to_static_leaf" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-2/paths-201/pathep-[eth1/1]"
}
`
