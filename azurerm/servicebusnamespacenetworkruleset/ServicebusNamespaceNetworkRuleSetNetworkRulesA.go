// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusnamespacenetworkruleset


type ServicebusNamespaceNetworkRuleSetNetworkRulesA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/servicebus_namespace_network_rule_set#subnet_id ServicebusNamespaceNetworkRuleSetA#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/servicebus_namespace_network_rule_set#ignore_missing_vnet_service_endpoint ServicebusNamespaceNetworkRuleSetA#ignore_missing_vnet_service_endpoint}.
	IgnoreMissingVnetServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVnetServiceEndpoint" yaml:"ignoreMissingVnetServiceEndpoint"`
}

