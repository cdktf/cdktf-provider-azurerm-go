// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#availability_set_name WorkloadsSapThreeTierVirtualInstance#availability_set_name}.
	AvailabilitySetName *string `field:"optional" json:"availabilitySetName" yaml:"availabilitySetName"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#load_balancer WorkloadsSapThreeTierVirtualInstance#load_balancer}
	LoadBalancer *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#virtual_machine WorkloadsSapThreeTierVirtualInstance#virtual_machine}
	VirtualMachine interface{} `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

