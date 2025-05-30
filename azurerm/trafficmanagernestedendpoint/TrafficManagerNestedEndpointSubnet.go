// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package trafficmanagernestedendpoint


type TrafficManagerNestedEndpointSubnet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/traffic_manager_nested_endpoint#first TrafficManagerNestedEndpoint#first}.
	First *string `field:"required" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/traffic_manager_nested_endpoint#last TrafficManagerNestedEndpoint#last}.
	Last *string `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/traffic_manager_nested_endpoint#scope TrafficManagerNestedEndpoint#scope}.
	Scope *float64 `field:"optional" json:"scope" yaml:"scope"`
}

