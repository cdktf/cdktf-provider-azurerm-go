// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusnamespacenetworkruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusNamespaceNetworkRuleSetAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#namespace_id ServicebusNamespaceNetworkRuleSetA#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#default_action ServicebusNamespaceNetworkRuleSetA#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#id ServicebusNamespaceNetworkRuleSetA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#ip_rules ServicebusNamespaceNetworkRuleSetA#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// network_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#network_rules ServicebusNamespaceNetworkRuleSetA#network_rules}
	NetworkRules interface{} `field:"optional" json:"networkRules" yaml:"networkRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#public_network_access_enabled ServicebusNamespaceNetworkRuleSetA#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#timeouts ServicebusNamespaceNetworkRuleSetA#timeouts}
	Timeouts *ServicebusNamespaceNetworkRuleSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/servicebus_namespace_network_rule_set#trusted_services_allowed ServicebusNamespaceNetworkRuleSetA#trusted_services_allowed}.
	TrustedServicesAllowed interface{} `field:"optional" json:"trustedServicesAllowed" yaml:"trustedServicesAllowed"`
}

