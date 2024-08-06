// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourcePimRouteMapEntryWithPimRouteMapPol(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPolAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "source_ip", "0.0.0.0"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPolAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPolAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.allow_test_2", "source_ip", "0.0.0.0"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "source_ip", "0.0.0.0"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigPimRouteMapEntryAllDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "action", "deny"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name", "name_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "source_ip", "1.1.1.1/30"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "order", "1"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigPimRouteMapEntryResetDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "source_ip", "0.0.0.0"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_pim_route_map_entry.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigPimRouteMapEntryChildrenDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "order", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "group_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "rendezvous_point_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_pim_route_map_entry.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigPimRouteMapEntryChildrenRemoveFromConfigDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigPimRouteMapEntryChildrenRemoveOneDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigPimRouteMapEntryChildrenRemoveAllDependencyWithPimRouteMapPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_pim_route_map_entry.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPolAllowExisting = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "allow_test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
}
resource "aci_pim_route_map_entry" "allow_test_2" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  depends_on = [aci_pim_route_map_entry.allow_test]
}
`

const testConfigPimRouteMapEntryMinDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
}
`

const testConfigPimRouteMapEntryAllDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  action = "deny"
  annotation = "annotation"
  description = "description_1"
  group_ip = "0.0.0.0"
  name = "name_1"
  name_alias = "name_alias_1"
  rendezvous_point_ip = "0.0.0.0"
  source_ip = "1.1.1.1/30"
}
`

const testConfigPimRouteMapEntryResetDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  action = "permit"
  annotation = "orchestrator:terraform"
  description = ""
  group_ip = "0.0.0.0"
  name = ""
  name_alias = ""
  rendezvous_point_ip = "0.0.0.0"
  source_ip = "0.0.0.0"
}
`
const testConfigPimRouteMapEntryChildrenDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigPimRouteMapEntryChildrenRemoveFromConfigDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
}
`

const testConfigPimRouteMapEntryChildrenRemoveOneDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigPimRouteMapEntryChildrenRemoveAllDependencyWithPimRouteMapPol = testConfigPimRouteMapPolMinDependencyWithFvTenant + `
resource "aci_pim_route_map_entry" "test" {
  parent_dn = aci_pim_route_map_policy.test.id
  order = "1"
  annotations = []
  tags = []
}
`
