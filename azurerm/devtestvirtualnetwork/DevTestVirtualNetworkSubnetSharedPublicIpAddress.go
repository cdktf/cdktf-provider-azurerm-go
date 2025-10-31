// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestvirtualnetwork


type DevTestVirtualNetworkSubnetSharedPublicIpAddress struct {
	// allowed_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/dev_test_virtual_network#allowed_ports DevTestVirtualNetwork#allowed_ports}
	AllowedPorts interface{} `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
}

