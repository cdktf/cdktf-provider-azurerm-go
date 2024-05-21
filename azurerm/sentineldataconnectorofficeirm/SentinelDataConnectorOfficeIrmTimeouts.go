// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorofficeirm


type SentinelDataConnectorOfficeIrmTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/sentinel_data_connector_office_irm#create SentinelDataConnectorOfficeIrm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/sentinel_data_connector_office_irm#delete SentinelDataConnectorOfficeIrm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/sentinel_data_connector_office_irm#read SentinelDataConnectorOfficeIrm#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

