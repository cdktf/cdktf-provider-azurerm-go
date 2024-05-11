// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/workloads_sap_three_tier_virtual_instance#availability_set_name WorkloadsSapThreeTierVirtualInstance#availability_set_name}.
	AvailabilitySetName *string `field:"optional" json:"availabilitySetName" yaml:"availabilitySetName"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/workloads_sap_three_tier_virtual_instance#virtual_machine WorkloadsSapThreeTierVirtualInstance#virtual_machine}
	VirtualMachine interface{} `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

