// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorazureadvancedthreatprotection


type SentinelDataConnectorAzureAdvancedThreatProtectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_advanced_threat_protection#create SentinelDataConnectorAzureAdvancedThreatProtection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_advanced_threat_protection#delete SentinelDataConnectorAzureAdvancedThreatProtection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sentinel_data_connector_azure_advanced_threat_protection#read SentinelDataConnectorAzureAdvancedThreatProtection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

