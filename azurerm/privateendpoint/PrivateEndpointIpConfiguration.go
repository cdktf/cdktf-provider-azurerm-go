// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privateendpoint


type PrivateEndpointIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/private_endpoint#name PrivateEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/private_endpoint#private_ip_address PrivateEndpoint#private_ip_address}.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/private_endpoint#member_name PrivateEndpoint#member_name}.
	MemberName *string `field:"optional" json:"memberName" yaml:"memberName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/private_endpoint#subresource_name PrivateEndpoint#subresource_name}.
	SubresourceName *string `field:"optional" json:"subresourceName" yaml:"subresourceName"`
}

