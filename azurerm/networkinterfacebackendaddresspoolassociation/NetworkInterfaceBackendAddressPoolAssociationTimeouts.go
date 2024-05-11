// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkinterfacebackendaddresspoolassociation


type NetworkInterfaceBackendAddressPoolAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_backend_address_pool_association#create NetworkInterfaceBackendAddressPoolAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_backend_address_pool_association#delete NetworkInterfaceBackendAddressPoolAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_backend_address_pool_association#read NetworkInterfaceBackendAddressPoolAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

