// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customprovider


type CustomProviderResourceType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/custom_provider#endpoint CustomProvider#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/custom_provider#name CustomProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/custom_provider#routing_type CustomProvider#routing_type}.
	RoutingType *string `field:"optional" json:"routingType" yaml:"routingType"`
}

