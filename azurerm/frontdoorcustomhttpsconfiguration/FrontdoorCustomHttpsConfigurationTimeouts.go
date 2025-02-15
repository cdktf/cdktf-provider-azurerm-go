// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorcustomhttpsconfiguration


type FrontdoorCustomHttpsConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/frontdoor_custom_https_configuration#create FrontdoorCustomHttpsConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/frontdoor_custom_https_configuration#delete FrontdoorCustomHttpsConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/frontdoor_custom_https_configuration#read FrontdoorCustomHttpsConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/frontdoor_custom_https_configuration#update FrontdoorCustomHttpsConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

