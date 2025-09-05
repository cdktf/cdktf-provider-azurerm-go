// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorazureactivedirectory


type SentinelDataConnectorAzureActiveDirectoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/sentinel_data_connector_azure_active_directory#create SentinelDataConnectorAzureActiveDirectory#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/sentinel_data_connector_azure_active_directory#delete SentinelDataConnectorAzureActiveDirectory#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/sentinel_data_connector_azure_active_directory#read SentinelDataConnectorAzureActiveDirectory#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

