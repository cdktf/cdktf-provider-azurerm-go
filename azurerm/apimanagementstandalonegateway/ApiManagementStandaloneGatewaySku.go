// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementstandalonegateway


type ApiManagementStandaloneGatewaySku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/api_management_standalone_gateway#name ApiManagementStandaloneGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/api_management_standalone_gateway#capacity ApiManagementStandaloneGateway#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

