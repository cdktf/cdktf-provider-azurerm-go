// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbbackendaddresspooladdress


type LbBackendAddressPoolAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/lb_backend_address_pool_address#create LbBackendAddressPoolAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/lb_backend_address_pool_address#delete LbBackendAddressPoolAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/lb_backend_address_pool_address#read LbBackendAddressPoolAddress#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/lb_backend_address_pool_address#update LbBackendAddressPoolAddress#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

