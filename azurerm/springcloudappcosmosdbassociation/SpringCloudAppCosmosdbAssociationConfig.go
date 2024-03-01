// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudappcosmosdbassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudAppCosmosdbAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#api_type SpringCloudAppCosmosdbAssociation#api_type}.
	ApiType *string `field:"required" json:"apiType" yaml:"apiType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_access_key SpringCloudAppCosmosdbAssociation#cosmosdb_access_key}.
	CosmosdbAccessKey *string `field:"required" json:"cosmosdbAccessKey" yaml:"cosmosdbAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_account_id SpringCloudAppCosmosdbAssociation#cosmosdb_account_id}.
	CosmosdbAccountId *string `field:"required" json:"cosmosdbAccountId" yaml:"cosmosdbAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#name SpringCloudAppCosmosdbAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#spring_cloud_app_id SpringCloudAppCosmosdbAssociation#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"required" json:"springCloudAppId" yaml:"springCloudAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_cassandra_keyspace_name SpringCloudAppCosmosdbAssociation#cosmosdb_cassandra_keyspace_name}.
	CosmosdbCassandraKeyspaceName *string `field:"optional" json:"cosmosdbCassandraKeyspaceName" yaml:"cosmosdbCassandraKeyspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_gremlin_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_gremlin_database_name}.
	CosmosdbGremlinDatabaseName *string `field:"optional" json:"cosmosdbGremlinDatabaseName" yaml:"cosmosdbGremlinDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_gremlin_graph_name SpringCloudAppCosmosdbAssociation#cosmosdb_gremlin_graph_name}.
	CosmosdbGremlinGraphName *string `field:"optional" json:"cosmosdbGremlinGraphName" yaml:"cosmosdbGremlinGraphName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_mongo_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_mongo_database_name}.
	CosmosdbMongoDatabaseName *string `field:"optional" json:"cosmosdbMongoDatabaseName" yaml:"cosmosdbMongoDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#cosmosdb_sql_database_name SpringCloudAppCosmosdbAssociation#cosmosdb_sql_database_name}.
	CosmosdbSqlDatabaseName *string `field:"optional" json:"cosmosdbSqlDatabaseName" yaml:"cosmosdbSqlDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#id SpringCloudAppCosmosdbAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_app_cosmosdb_association#timeouts SpringCloudAppCosmosdbAssociation#timeouts}
	Timeouts *SpringCloudAppCosmosdbAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

