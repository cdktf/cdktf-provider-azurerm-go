// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfigurationpolicygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnServerConfigurationPolicyGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#name VpnServerConfigurationPolicyGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#policy VpnServerConfigurationPolicyGroup#policy}
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#vpn_server_configuration_id VpnServerConfigurationPolicyGroup#vpn_server_configuration_id}.
	VpnServerConfigurationId *string `field:"required" json:"vpnServerConfigurationId" yaml:"vpnServerConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#id VpnServerConfigurationPolicyGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#is_default VpnServerConfigurationPolicyGroup#is_default}.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#priority VpnServerConfigurationPolicyGroup#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/vpn_server_configuration_policy_group#timeouts VpnServerConfigurationPolicyGroup#timeouts}
	Timeouts *VpnServerConfigurationPolicyGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

