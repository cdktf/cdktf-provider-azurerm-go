// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqldedicatedgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbSqlDedicatedGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/cosmosdb_sql_dedicated_gateway#cosmosdb_account_id CosmosdbSqlDedicatedGateway#cosmosdb_account_id}.
	CosmosdbAccountId *string `field:"required" json:"cosmosdbAccountId" yaml:"cosmosdbAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/cosmosdb_sql_dedicated_gateway#instance_count CosmosdbSqlDedicatedGateway#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/cosmosdb_sql_dedicated_gateway#instance_size CosmosdbSqlDedicatedGateway#instance_size}.
	InstanceSize *string `field:"required" json:"instanceSize" yaml:"instanceSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/cosmosdb_sql_dedicated_gateway#id CosmosdbSqlDedicatedGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/cosmosdb_sql_dedicated_gateway#timeouts CosmosdbSqlDedicatedGateway#timeouts}
	Timeouts *CosmosdbSqlDedicatedGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

