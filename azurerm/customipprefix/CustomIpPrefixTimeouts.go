// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customipprefix


type CustomIpPrefixTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/custom_ip_prefix#create CustomIpPrefix#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/custom_ip_prefix#delete CustomIpPrefix#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/custom_ip_prefix#read CustomIpPrefix#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/custom_ip_prefix#update CustomIpPrefix#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

