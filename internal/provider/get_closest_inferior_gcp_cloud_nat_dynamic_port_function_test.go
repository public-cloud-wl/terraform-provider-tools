// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestGetClosestInferiorGCPCloudNatDynamicPortFunction_Equals(t *testing.T) {
	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccDefaultFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				output "test" {
					value = provider::tools::GetClosestInferiorGCPCloudNatDynamicPortFunction("min", 8192)
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "8192"),
				),
			},
		},
	})
}

func TestGetClosestInferiorGCPCloudNatDynamicPortFunction_MoreMin(t *testing.T) {
	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccDefaultFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				output "test" {
					value = provider::tools::GetClosestInferiorGCPCloudNatDynamicPortFunction("min", 65537)
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "32768"),
				),
			},
		},
	})
}

func TestGetClosestInferiorGCPCloudNatDynamicPortFunction_MoreMax(t *testing.T) {
	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccDefaultFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				output "test" {
					value = provider::tools::GetClosestInferiorGCPCloudNatDynamicPortFunction("max", 65537)
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "65536"),
				),
			},
		},
	})
}

func TestGetClosestInferiorGCPCloudNatDynamicPortFunction_Error(t *testing.T) {
	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccDefaultFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				output "test" {
					value = provider::tools::GetClosestInferiorGCPCloudNatDynamicPortFunction("toto", 42)
				}
				`,
				// The parameter does not enable AllowNullValue
				ExpectError: regexp.MustCompile(`Invalid port_type`),
			},
		},
	})
}
