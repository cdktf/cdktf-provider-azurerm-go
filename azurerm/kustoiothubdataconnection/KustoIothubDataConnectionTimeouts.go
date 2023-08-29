// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoiothubdataconnection


type KustoIothubDataConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kusto_iothub_data_connection#create KustoIothubDataConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kusto_iothub_data_connection#delete KustoIothubDataConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kusto_iothub_data_connection#read KustoIothubDataConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

