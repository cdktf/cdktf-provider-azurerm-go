// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkinterfaceapplicationgatewaybackendaddresspoolassociation


type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_interface_application_gateway_backend_address_pool_association#create NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_interface_application_gateway_backend_address_pool_association#delete NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_interface_application_gateway_backend_address_pool_association#read NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

