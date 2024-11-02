// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksimgroup


type MobileNetworkSimGroupIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mobile_network_sim_group#identity_ids MobileNetworkSimGroup#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mobile_network_sim_group#type MobileNetworkSimGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

