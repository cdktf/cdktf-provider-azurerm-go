// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestvirtualnetwork


type DevTestVirtualNetworkSubnet struct {
	// shared_public_ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/dev_test_virtual_network#shared_public_ip_address DevTestVirtualNetwork#shared_public_ip_address}
	SharedPublicIpAddress *DevTestVirtualNetworkSubnetSharedPublicIpAddress `field:"optional" json:"sharedPublicIpAddress" yaml:"sharedPublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/dev_test_virtual_network#use_in_virtual_machine_creation DevTestVirtualNetwork#use_in_virtual_machine_creation}.
	UseInVirtualMachineCreation *string `field:"optional" json:"useInVirtualMachineCreation" yaml:"useInVirtualMachineCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/dev_test_virtual_network#use_public_ip_address DevTestVirtualNetwork#use_public_ip_address}.
	UsePublicIpAddress *string `field:"optional" json:"usePublicIpAddress" yaml:"usePublicIpAddress"`
}

