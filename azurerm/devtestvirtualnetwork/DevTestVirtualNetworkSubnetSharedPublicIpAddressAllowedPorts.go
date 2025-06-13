// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestvirtualnetwork


type DevTestVirtualNetworkSubnetSharedPublicIpAddressAllowedPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/dev_test_virtual_network#backend_port DevTestVirtualNetwork#backend_port}.
	BackendPort *float64 `field:"optional" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/dev_test_virtual_network#transport_protocol DevTestVirtualNetwork#transport_protocol}.
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
}

