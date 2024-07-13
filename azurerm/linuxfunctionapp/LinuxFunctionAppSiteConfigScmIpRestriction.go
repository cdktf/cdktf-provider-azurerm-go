// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppSiteConfigScmIpRestriction struct {
	// The action to take. Possible values are `Allow` or `Deny`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#action LinuxFunctionApp#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The description of the IP restriction rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#description LinuxFunctionApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#headers LinuxFunctionApp#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The CIDR notation of the IP or IP Range to match.
	//
	// For example: `10.0.0.0/24` or `192.168.10.1/32` or `fe80::/64` or `13.107.6.152/31,13.107.128.0/22`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#ip_address LinuxFunctionApp#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The name which should be used for this `ip_restriction`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#name LinuxFunctionApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The priority value of this `ip_restriction`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#priority LinuxFunctionApp#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The Service Tag used for this IP Restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#service_tag LinuxFunctionApp#service_tag}
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// The Virtual Network Subnet ID used for this IP Restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/linux_function_app#virtual_network_subnet_id LinuxFunctionApp#virtual_network_subnet_id}
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

