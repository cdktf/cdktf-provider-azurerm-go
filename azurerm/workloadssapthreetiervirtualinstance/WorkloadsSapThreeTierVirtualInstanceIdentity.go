// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/workloads_sap_three_tier_virtual_instance#identity_ids WorkloadsSapThreeTierVirtualInstance#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/workloads_sap_three_tier_virtual_instance#type WorkloadsSapThreeTierVirtualInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

