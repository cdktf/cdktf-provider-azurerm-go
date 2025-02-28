// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry


type ContainerRegistryNetworkRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_registry#default_action ContainerRegistry#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_registry#ip_rule ContainerRegistry#ip_rule}.
	IpRule interface{} `field:"optional" json:"ipRule" yaml:"ipRule"`
}

