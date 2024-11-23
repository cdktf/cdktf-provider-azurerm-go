// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputcosmosdb

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputCosmosdbConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#container_name StreamAnalyticsOutputCosmosdb#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#cosmosdb_account_key StreamAnalyticsOutputCosmosdb#cosmosdb_account_key}.
	CosmosdbAccountKey *string `field:"required" json:"cosmosdbAccountKey" yaml:"cosmosdbAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#cosmosdb_sql_database_id StreamAnalyticsOutputCosmosdb#cosmosdb_sql_database_id}.
	CosmosdbSqlDatabaseId *string `field:"required" json:"cosmosdbSqlDatabaseId" yaml:"cosmosdbSqlDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#name StreamAnalyticsOutputCosmosdb#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#stream_analytics_job_id StreamAnalyticsOutputCosmosdb#stream_analytics_job_id}.
	StreamAnalyticsJobId *string `field:"required" json:"streamAnalyticsJobId" yaml:"streamAnalyticsJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#document_id StreamAnalyticsOutputCosmosdb#document_id}.
	DocumentId *string `field:"optional" json:"documentId" yaml:"documentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#id StreamAnalyticsOutputCosmosdb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#partition_key StreamAnalyticsOutputCosmosdb#partition_key}.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stream_analytics_output_cosmosdb#timeouts StreamAnalyticsOutputCosmosdb#timeouts}
	Timeouts *StreamAnalyticsOutputCosmosdbTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

