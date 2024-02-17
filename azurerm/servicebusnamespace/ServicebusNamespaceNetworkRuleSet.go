// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusnamespace


type ServicebusNamespaceNetworkRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/servicebus_namespace#default_action ServicebusNamespace#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/servicebus_namespace#ip_rules ServicebusNamespace#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// network_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/servicebus_namespace#network_rules ServicebusNamespace#network_rules}
	NetworkRules interface{} `field:"optional" json:"networkRules" yaml:"networkRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/servicebus_namespace#public_network_access_enabled ServicebusNamespace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/servicebus_namespace#trusted_services_allowed ServicebusNamespace#trusted_services_allowed}.
	TrustedServicesAllowed interface{} `field:"optional" json:"trustedServicesAllowed" yaml:"trustedServicesAllowed"`
}

