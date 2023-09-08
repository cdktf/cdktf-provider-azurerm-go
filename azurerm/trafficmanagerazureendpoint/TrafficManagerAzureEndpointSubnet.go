// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerazureendpoint


type TrafficManagerAzureEndpointSubnet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/traffic_manager_azure_endpoint#first TrafficManagerAzureEndpoint#first}.
	First *string `field:"required" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/traffic_manager_azure_endpoint#last TrafficManagerAzureEndpoint#last}.
	Last *string `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/traffic_manager_azure_endpoint#scope TrafficManagerAzureEndpoint#scope}.
	Scope *float64 `field:"optional" json:"scope" yaml:"scope"`
}

