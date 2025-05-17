// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoeventhubdataconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoEventhubDataConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#cluster_name KustoEventhubDataConnection#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#consumer_group KustoEventhubDataConnection#consumer_group}.
	ConsumerGroup *string `field:"required" json:"consumerGroup" yaml:"consumerGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#database_name KustoEventhubDataConnection#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#eventhub_id KustoEventhubDataConnection#eventhub_id}.
	EventhubId *string `field:"required" json:"eventhubId" yaml:"eventhubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#location KustoEventhubDataConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#name KustoEventhubDataConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#resource_group_name KustoEventhubDataConnection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#compression KustoEventhubDataConnection#compression}.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#database_routing_type KustoEventhubDataConnection#database_routing_type}.
	DatabaseRoutingType *string `field:"optional" json:"databaseRoutingType" yaml:"databaseRoutingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#data_format KustoEventhubDataConnection#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#event_system_properties KustoEventhubDataConnection#event_system_properties}.
	EventSystemProperties *[]*string `field:"optional" json:"eventSystemProperties" yaml:"eventSystemProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#id KustoEventhubDataConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#identity_id KustoEventhubDataConnection#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#mapping_rule_name KustoEventhubDataConnection#mapping_rule_name}.
	MappingRuleName *string `field:"optional" json:"mappingRuleName" yaml:"mappingRuleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#table_name KustoEventhubDataConnection#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kusto_eventhub_data_connection#timeouts KustoEventhubDataConnection#timeouts}
	Timeouts *KustoEventhubDataConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

