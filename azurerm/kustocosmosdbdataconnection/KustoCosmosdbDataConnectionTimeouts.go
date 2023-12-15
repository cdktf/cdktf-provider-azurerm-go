// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustocosmosdbdataconnection


type KustoCosmosdbDataConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/kusto_cosmosdb_data_connection#create KustoCosmosdbDataConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/kusto_cosmosdb_data_connection#delete KustoCosmosdbDataConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/kusto_cosmosdb_data_connection#read KustoCosmosdbDataConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

