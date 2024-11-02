// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusnamespace


type ServicebusNamespaceNetworkRuleSetNetworkRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/servicebus_namespace#subnet_id ServicebusNamespace#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/servicebus_namespace#ignore_missing_vnet_service_endpoint ServicebusNamespace#ignore_missing_vnet_service_endpoint}.
	IgnoreMissingVnetServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVnetServiceEndpoint" yaml:"ignoreMissingVnetServiceEndpoint"`
}

