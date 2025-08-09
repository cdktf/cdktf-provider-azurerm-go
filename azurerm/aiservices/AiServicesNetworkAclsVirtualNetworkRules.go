// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aiservices


type AiServicesNetworkAclsVirtualNetworkRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/ai_services#subnet_id AiServices#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/ai_services#ignore_missing_vnet_service_endpoint AiServices#ignore_missing_vnet_service_endpoint}.
	IgnoreMissingVnetServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVnetServiceEndpoint" yaml:"ignoreMissingVnetServiceEndpoint"`
}

