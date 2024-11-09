// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoeventgriddataconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoEventgridDataConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#cluster_name KustoEventgridDataConnection#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#database_name KustoEventgridDataConnection#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#eventhub_consumer_group_name KustoEventgridDataConnection#eventhub_consumer_group_name}.
	EventhubConsumerGroupName *string `field:"required" json:"eventhubConsumerGroupName" yaml:"eventhubConsumerGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#eventhub_id KustoEventgridDataConnection#eventhub_id}.
	EventhubId *string `field:"required" json:"eventhubId" yaml:"eventhubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#location KustoEventgridDataConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#name KustoEventgridDataConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#resource_group_name KustoEventgridDataConnection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#storage_account_id KustoEventgridDataConnection#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#blob_storage_event_type KustoEventgridDataConnection#blob_storage_event_type}.
	BlobStorageEventType *string `field:"optional" json:"blobStorageEventType" yaml:"blobStorageEventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#database_routing_type KustoEventgridDataConnection#database_routing_type}.
	DatabaseRoutingType *string `field:"optional" json:"databaseRoutingType" yaml:"databaseRoutingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#data_format KustoEventgridDataConnection#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#eventgrid_resource_id KustoEventgridDataConnection#eventgrid_resource_id}.
	EventgridResourceId *string `field:"optional" json:"eventgridResourceId" yaml:"eventgridResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#id KustoEventgridDataConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#managed_identity_resource_id KustoEventgridDataConnection#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"optional" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#mapping_rule_name KustoEventgridDataConnection#mapping_rule_name}.
	MappingRuleName *string `field:"optional" json:"mappingRuleName" yaml:"mappingRuleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#skip_first_record KustoEventgridDataConnection#skip_first_record}.
	SkipFirstRecord interface{} `field:"optional" json:"skipFirstRecord" yaml:"skipFirstRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#table_name KustoEventgridDataConnection#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kusto_eventgrid_data_connection#timeouts KustoEventgridDataConnection#timeouts}
	Timeouts *KustoEventgridDataConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

