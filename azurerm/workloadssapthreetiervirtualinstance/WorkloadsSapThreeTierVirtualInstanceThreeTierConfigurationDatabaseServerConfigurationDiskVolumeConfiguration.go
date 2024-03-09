// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationDiskVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/workloads_sap_three_tier_virtual_instance#number_of_disks WorkloadsSapThreeTierVirtualInstance#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/workloads_sap_three_tier_virtual_instance#size_in_gb WorkloadsSapThreeTierVirtualInstance#size_in_gb}.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/workloads_sap_three_tier_virtual_instance#sku_name WorkloadsSapThreeTierVirtualInstance#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/workloads_sap_three_tier_virtual_instance#volume_name WorkloadsSapThreeTierVirtualInstance#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
}

