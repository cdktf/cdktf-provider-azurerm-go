// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationVirtualMachineConfiguration struct {
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_three_tier_virtual_instance#image WorkloadsSapThreeTierVirtualInstance#image}
	Image *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationVirtualMachineConfigurationImage `field:"required" json:"image" yaml:"image"`
	// os_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_three_tier_virtual_instance#os_profile WorkloadsSapThreeTierVirtualInstance#os_profile}
	OsProfile *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationVirtualMachineConfigurationOsProfile `field:"required" json:"osProfile" yaml:"osProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_three_tier_virtual_instance#virtual_machine_size WorkloadsSapThreeTierVirtualInstance#virtual_machine_size}.
	VirtualMachineSize *string `field:"required" json:"virtualMachineSize" yaml:"virtualMachineSize"`
}

