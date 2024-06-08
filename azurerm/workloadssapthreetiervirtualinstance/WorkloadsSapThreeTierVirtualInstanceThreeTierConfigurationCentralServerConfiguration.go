// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/workloads_sap_three_tier_virtual_instance#instance_count WorkloadsSapThreeTierVirtualInstance#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/workloads_sap_three_tier_virtual_instance#subnet_id WorkloadsSapThreeTierVirtualInstance#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// virtual_machine_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/workloads_sap_three_tier_virtual_instance#virtual_machine_configuration WorkloadsSapThreeTierVirtualInstance#virtual_machine_configuration}
	VirtualMachineConfiguration *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfigurationVirtualMachineConfiguration `field:"required" json:"virtualMachineConfiguration" yaml:"virtualMachineConfiguration"`
}

