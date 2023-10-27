// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryhypervnetworkmapping


type SiteRecoveryHypervNetworkMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/site_recovery_hyperv_network_mapping#create SiteRecoveryHypervNetworkMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/site_recovery_hyperv_network_mapping#delete SiteRecoveryHypervNetworkMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/site_recovery_hyperv_network_mapping#read SiteRecoveryHypervNetworkMapping#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

