// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolNetworkConfigurationEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/batch_pool#backend_port BatchPool#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/batch_pool#frontend_port_range BatchPool#frontend_port_range}.
	FrontendPortRange *string `field:"required" json:"frontendPortRange" yaml:"frontendPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/batch_pool#protocol BatchPool#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// network_security_group_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/batch_pool#network_security_group_rules BatchPool#network_security_group_rules}
	NetworkSecurityGroupRules interface{} `field:"optional" json:"networkSecurityGroupRules" yaml:"networkSecurityGroupRules"`
}

