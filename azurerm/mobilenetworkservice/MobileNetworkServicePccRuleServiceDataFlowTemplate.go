// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice


type MobileNetworkServicePccRuleServiceDataFlowTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mobile_network_service#direction MobileNetworkService#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mobile_network_service#name MobileNetworkService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mobile_network_service#protocol MobileNetworkService#protocol}.
	Protocol *[]*string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mobile_network_service#remote_ip_list MobileNetworkService#remote_ip_list}.
	RemoteIpList *[]*string `field:"required" json:"remoteIpList" yaml:"remoteIpList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mobile_network_service#ports MobileNetworkService#ports}.
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
}

