// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanageripampoolstaticcidr


type NetworkManagerIpamPoolStaticCidrTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_ipam_pool_static_cidr#create NetworkManagerIpamPoolStaticCidr#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_ipam_pool_static_cidr#delete NetworkManagerIpamPoolStaticCidr#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_ipam_pool_static_cidr#read NetworkManagerIpamPoolStaticCidr#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_ipam_pool_static_cidr#update NetworkManagerIpamPoolStaticCidr#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

