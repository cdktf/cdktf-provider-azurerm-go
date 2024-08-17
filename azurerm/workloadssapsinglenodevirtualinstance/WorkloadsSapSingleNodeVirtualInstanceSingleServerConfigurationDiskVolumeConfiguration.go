// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance


type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationDiskVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_single_node_virtual_instance#number_of_disks WorkloadsSapSingleNodeVirtualInstance#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_single_node_virtual_instance#size_in_gb WorkloadsSapSingleNodeVirtualInstance#size_in_gb}.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_single_node_virtual_instance#sku_name WorkloadsSapSingleNodeVirtualInstance#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/workloads_sap_single_node_virtual_instance#volume_name WorkloadsSapSingleNodeVirtualInstance#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
}

