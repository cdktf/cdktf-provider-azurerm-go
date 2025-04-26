// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagerazureendpoint


type TrafficManagerAzureEndpointCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/traffic_manager_azure_endpoint#name TrafficManagerAzureEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/traffic_manager_azure_endpoint#value TrafficManagerAzureEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

