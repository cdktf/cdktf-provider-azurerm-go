// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustocosmosdbdataconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoCosmosdbDataConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#cosmosdb_container_id KustoCosmosdbDataConnection#cosmosdb_container_id}.
	CosmosdbContainerId *string `field:"required" json:"cosmosdbContainerId" yaml:"cosmosdbContainerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#kusto_database_id KustoCosmosdbDataConnection#kusto_database_id}.
	KustoDatabaseId *string `field:"required" json:"kustoDatabaseId" yaml:"kustoDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#location KustoCosmosdbDataConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#managed_identity_id KustoCosmosdbDataConnection#managed_identity_id}.
	ManagedIdentityId *string `field:"required" json:"managedIdentityId" yaml:"managedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#name KustoCosmosdbDataConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#table_name KustoCosmosdbDataConnection#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#id KustoCosmosdbDataConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#mapping_rule_name KustoCosmosdbDataConnection#mapping_rule_name}.
	MappingRuleName *string `field:"optional" json:"mappingRuleName" yaml:"mappingRuleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#retrieval_start_date KustoCosmosdbDataConnection#retrieval_start_date}.
	RetrievalStartDate *string `field:"optional" json:"retrievalStartDate" yaml:"retrievalStartDate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/kusto_cosmosdb_data_connection#timeouts KustoCosmosdbDataConnection#timeouts}
	Timeouts *KustoCosmosdbDataConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

