// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentineldataconnectoriot


type SentinelDataConnectorIotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_data_connector_iot#create SentinelDataConnectorIot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_data_connector_iot#delete SentinelDataConnectorIot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/sentinel_data_connector_iot#read SentinelDataConnectorIot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

