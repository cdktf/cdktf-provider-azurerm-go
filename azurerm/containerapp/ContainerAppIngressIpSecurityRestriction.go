// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppIngressIpSecurityRestriction struct {
	// The action. Allow or Deny.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/container_app#action ContainerApp#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The incoming IP address or range of IP addresses (in CIDR notation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/container_app#ip_address_range ContainerApp#ip_address_range}
	IpAddressRange *string `field:"required" json:"ipAddressRange" yaml:"ipAddressRange"`
	// Name for the IP restriction rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describe the IP restriction rule that is being sent to the container-app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/container_app#description ContainerApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

