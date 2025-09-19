// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewaynatrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayNatRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#name VpnGatewayNatRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#vpn_gateway_id VpnGatewayNatRule#vpn_gateway_id}.
	VpnGatewayId *string `field:"required" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// external_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#external_mapping VpnGatewayNatRule#external_mapping}
	ExternalMapping interface{} `field:"optional" json:"externalMapping" yaml:"externalMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#id VpnGatewayNatRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// internal_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#internal_mapping VpnGatewayNatRule#internal_mapping}
	InternalMapping interface{} `field:"optional" json:"internalMapping" yaml:"internalMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#ip_configuration_id VpnGatewayNatRule#ip_configuration_id}.
	IpConfigurationId *string `field:"optional" json:"ipConfigurationId" yaml:"ipConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#mode VpnGatewayNatRule#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#timeouts VpnGatewayNatRule#timeouts}
	Timeouts *VpnGatewayNatRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/vpn_gateway_nat_rule#type VpnGatewayNatRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

