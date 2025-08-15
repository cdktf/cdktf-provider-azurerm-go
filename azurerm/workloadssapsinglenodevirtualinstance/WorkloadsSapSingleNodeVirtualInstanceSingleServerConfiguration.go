// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance


type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#app_resource_group_name WorkloadsSapSingleNodeVirtualInstance#app_resource_group_name}.
	AppResourceGroupName *string `field:"required" json:"appResourceGroupName" yaml:"appResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#subnet_id WorkloadsSapSingleNodeVirtualInstance#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// virtual_machine_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#virtual_machine_configuration WorkloadsSapSingleNodeVirtualInstance#virtual_machine_configuration}
	VirtualMachineConfiguration *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration `field:"required" json:"virtualMachineConfiguration" yaml:"virtualMachineConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#database_type WorkloadsSapSingleNodeVirtualInstance#database_type}.
	DatabaseType *string `field:"optional" json:"databaseType" yaml:"databaseType"`
	// disk_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#disk_volume_configuration WorkloadsSapSingleNodeVirtualInstance#disk_volume_configuration}
	DiskVolumeConfiguration interface{} `field:"optional" json:"diskVolumeConfiguration" yaml:"diskVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#secondary_ip_enabled WorkloadsSapSingleNodeVirtualInstance#secondary_ip_enabled}.
	SecondaryIpEnabled interface{} `field:"optional" json:"secondaryIpEnabled" yaml:"secondaryIpEnabled"`
	// virtual_machine_resource_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/workloads_sap_single_node_virtual_instance#virtual_machine_resource_names WorkloadsSapSingleNodeVirtualInstance#virtual_machine_resource_names}
	VirtualMachineResourceNames *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames `field:"optional" json:"virtualMachineResourceNames" yaml:"virtualMachineResourceNames"`
}

