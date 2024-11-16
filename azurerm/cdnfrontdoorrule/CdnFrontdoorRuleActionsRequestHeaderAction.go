// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActionsRequestHeaderAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/cdn_frontdoor_rule#header_action CdnFrontdoorRule#header_action}.
	HeaderAction *string `field:"required" json:"headerAction" yaml:"headerAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/cdn_frontdoor_rule#header_name CdnFrontdoorRule#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/cdn_frontdoor_rule#value CdnFrontdoorRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

