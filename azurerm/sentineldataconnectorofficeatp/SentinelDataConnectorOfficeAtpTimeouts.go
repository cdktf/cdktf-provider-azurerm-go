// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectorofficeatp


type SentinelDataConnectorOfficeAtpTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_data_connector_office_atp#create SentinelDataConnectorOfficeAtp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_data_connector_office_atp#delete SentinelDataConnectorOfficeAtp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_data_connector_office_atp#read SentinelDataConnectorOfficeAtp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

