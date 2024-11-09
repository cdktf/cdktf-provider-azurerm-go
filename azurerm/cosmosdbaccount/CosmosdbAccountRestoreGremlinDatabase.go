// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountRestoreGremlinDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_account#graph_names CosmosdbAccount#graph_names}.
	GraphNames *[]*string `field:"optional" json:"graphNames" yaml:"graphNames"`
}

