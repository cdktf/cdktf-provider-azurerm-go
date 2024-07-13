// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewaynatrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkGatewayNatRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// external_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#external_mapping VirtualNetworkGatewayNatRule#external_mapping}
	ExternalMapping interface{} `field:"required" json:"externalMapping" yaml:"externalMapping"`
	// internal_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#internal_mapping VirtualNetworkGatewayNatRule#internal_mapping}
	InternalMapping interface{} `field:"required" json:"internalMapping" yaml:"internalMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#name VirtualNetworkGatewayNatRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#resource_group_name VirtualNetworkGatewayNatRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#virtual_network_gateway_id VirtualNetworkGatewayNatRule#virtual_network_gateway_id}.
	VirtualNetworkGatewayId *string `field:"required" json:"virtualNetworkGatewayId" yaml:"virtualNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#id VirtualNetworkGatewayNatRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#ip_configuration_id VirtualNetworkGatewayNatRule#ip_configuration_id}.
	IpConfigurationId *string `field:"optional" json:"ipConfigurationId" yaml:"ipConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#mode VirtualNetworkGatewayNatRule#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#timeouts VirtualNetworkGatewayNatRule#timeouts}
	Timeouts *VirtualNetworkGatewayNatRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/virtual_network_gateway_nat_rule#type VirtualNetworkGatewayNatRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

