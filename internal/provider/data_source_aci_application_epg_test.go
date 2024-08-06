// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvAEPgWithFvAp(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvAEPgDataSourceDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "contract_exception_tag", ""),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "flood_in_encapsulation", "disabled"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "forwarding_control", "none"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "has_multicast_source", "no"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "intra_epg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "priority", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_application_epg.test", "useg_epg", "no"),
					resource.TestCheckResourceAttrSet("data.aci_application_epg.test", "pc_tag"),
				),
			},
			{
				Config:      testConfigFvAEPgNotExistingFvAp,
				ExpectError: regexp.MustCompile("Failed to read aci_application_epg data source"),
			},
		},
	})
}

const testConfigFvAEPgDataSourceDependencyWithFvAp = testConfigFvAEPgMinDependencyWithFvAp + `
data "aci_application_epg" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  depends_on = [aci_application_epg.test]
}
`

const testConfigFvAEPgNotExistingFvAp = testConfigFvApMinDependencyWithFvTenant + `
data "aci_application_epg" "test_non_existing" {
  parent_dn = aci_application_profile.test.id
  name = "non_existing_name"
}
`
