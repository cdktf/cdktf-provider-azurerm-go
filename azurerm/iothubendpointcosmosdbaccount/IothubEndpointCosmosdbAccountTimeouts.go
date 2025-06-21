// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointcosmosdbaccount


type IothubEndpointCosmosdbAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_cosmosdb_account#create IothubEndpointCosmosdbAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_cosmosdb_account#delete IothubEndpointCosmosdbAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_cosmosdb_account#read IothubEndpointCosmosdbAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_cosmosdb_account#update IothubEndpointCosmosdbAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

