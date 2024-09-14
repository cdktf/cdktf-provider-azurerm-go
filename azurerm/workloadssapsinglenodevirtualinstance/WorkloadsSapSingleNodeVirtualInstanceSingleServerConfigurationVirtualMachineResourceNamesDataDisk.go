// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance


type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNamesDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/workloads_sap_single_node_virtual_instance#names WorkloadsSapSingleNodeVirtualInstance#names}.
	Names *[]*string `field:"required" json:"names" yaml:"names"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/workloads_sap_single_node_virtual_instance#volume_name WorkloadsSapSingleNodeVirtualInstance#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
}

