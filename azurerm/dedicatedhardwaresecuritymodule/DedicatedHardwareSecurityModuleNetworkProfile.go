// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dedicatedhardwaresecuritymodule


type DedicatedHardwareSecurityModuleNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/dedicated_hardware_security_module#network_interface_private_ip_addresses DedicatedHardwareSecurityModule#network_interface_private_ip_addresses}.
	NetworkInterfacePrivateIpAddresses *[]*string `field:"required" json:"networkInterfacePrivateIpAddresses" yaml:"networkInterfacePrivateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/dedicated_hardware_security_module#subnet_id DedicatedHardwareSecurityModule#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

