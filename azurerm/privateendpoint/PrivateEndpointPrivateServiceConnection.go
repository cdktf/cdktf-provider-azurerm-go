// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privateendpoint


type PrivateEndpointPrivateServiceConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#is_manual_connection PrivateEndpoint#is_manual_connection}.
	IsManualConnection interface{} `field:"required" json:"isManualConnection" yaml:"isManualConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#name PrivateEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#private_connection_resource_alias PrivateEndpoint#private_connection_resource_alias}.
	PrivateConnectionResourceAlias *string `field:"optional" json:"privateConnectionResourceAlias" yaml:"privateConnectionResourceAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#private_connection_resource_id PrivateEndpoint#private_connection_resource_id}.
	PrivateConnectionResourceId *string `field:"optional" json:"privateConnectionResourceId" yaml:"privateConnectionResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#request_message PrivateEndpoint#request_message}.
	RequestMessage *string `field:"optional" json:"requestMessage" yaml:"requestMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/private_endpoint#subresource_names PrivateEndpoint#subresource_names}.
	SubresourceNames *[]*string `field:"optional" json:"subresourceNames" yaml:"subresourceNames"`
}

