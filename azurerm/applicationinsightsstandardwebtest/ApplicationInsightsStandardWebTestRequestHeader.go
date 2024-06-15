// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestRequestHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/application_insights_standard_web_test#name ApplicationInsightsStandardWebTest#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/application_insights_standard_web_test#value ApplicationInsightsStandardWebTest#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

