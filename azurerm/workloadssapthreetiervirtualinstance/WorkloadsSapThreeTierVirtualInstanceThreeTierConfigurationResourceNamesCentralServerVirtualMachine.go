// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerVirtualMachine struct {
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/workloads_sap_three_tier_virtual_instance#data_disk WorkloadsSapThreeTierVirtualInstance#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/workloads_sap_three_tier_virtual_instance#host_name WorkloadsSapThreeTierVirtualInstance#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/workloads_sap_three_tier_virtual_instance#network_interface_names WorkloadsSapThreeTierVirtualInstance#network_interface_names}.
	NetworkInterfaceNames *[]*string `field:"optional" json:"networkInterfaceNames" yaml:"networkInterfaceNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/workloads_sap_three_tier_virtual_instance#os_disk_name WorkloadsSapThreeTierVirtualInstance#os_disk_name}.
	OsDiskName *string `field:"optional" json:"osDiskName" yaml:"osDiskName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/workloads_sap_three_tier_virtual_instance#virtual_machine_name WorkloadsSapThreeTierVirtualInstance#virtual_machine_name}.
	VirtualMachineName *string `field:"optional" json:"virtualMachineName" yaml:"virtualMachineName"`
}

