// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfigurationImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/workloads_sap_three_tier_virtual_instance#offer WorkloadsSapThreeTierVirtualInstance#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/workloads_sap_three_tier_virtual_instance#publisher WorkloadsSapThreeTierVirtualInstance#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/workloads_sap_three_tier_virtual_instance#sku WorkloadsSapThreeTierVirtualInstance#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/workloads_sap_three_tier_virtual_instance#version WorkloadsSapThreeTierVirtualInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

