// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance


type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration struct {
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/workloads_sap_single_node_virtual_instance#image WorkloadsSapSingleNodeVirtualInstance#image}
	Image *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationImage `field:"required" json:"image" yaml:"image"`
	// os_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/workloads_sap_single_node_virtual_instance#os_profile WorkloadsSapSingleNodeVirtualInstance#os_profile}
	OsProfile *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationOsProfile `field:"required" json:"osProfile" yaml:"osProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/workloads_sap_single_node_virtual_instance#virtual_machine_size WorkloadsSapSingleNodeVirtualInstance#virtual_machine_size}.
	VirtualMachineSize *string `field:"required" json:"virtualMachineSize" yaml:"virtualMachineSize"`
}

