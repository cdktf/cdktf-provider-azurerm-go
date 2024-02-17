// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsstandardwebtest


type ApplicationInsightsStandardWebTestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/application_insights_standard_web_test#create ApplicationInsightsStandardWebTest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/application_insights_standard_web_test#delete ApplicationInsightsStandardWebTest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/application_insights_standard_web_test#read ApplicationInsightsStandardWebTest#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/application_insights_standard_web_test#update ApplicationInsightsStandardWebTest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

