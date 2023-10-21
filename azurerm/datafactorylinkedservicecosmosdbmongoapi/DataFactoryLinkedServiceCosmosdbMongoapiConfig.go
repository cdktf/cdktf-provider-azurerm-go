// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicecosmosdbmongoapi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceCosmosdbMongoapiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#data_factory_id DataFactoryLinkedServiceCosmosdbMongoapi#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#name DataFactoryLinkedServiceCosmosdbMongoapi#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#additional_properties DataFactoryLinkedServiceCosmosdbMongoapi#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#annotations DataFactoryLinkedServiceCosmosdbMongoapi#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#connection_string DataFactoryLinkedServiceCosmosdbMongoapi#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#database DataFactoryLinkedServiceCosmosdbMongoapi#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#description DataFactoryLinkedServiceCosmosdbMongoapi#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#id DataFactoryLinkedServiceCosmosdbMongoapi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#integration_runtime_name DataFactoryLinkedServiceCosmosdbMongoapi#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#parameters DataFactoryLinkedServiceCosmosdbMongoapi#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#server_version_is_32_or_higher DataFactoryLinkedServiceCosmosdbMongoapi#server_version_is_32_or_higher}.
	ServerVersionIs32OrHigher interface{} `field:"optional" json:"serverVersionIs32OrHigher" yaml:"serverVersionIs32OrHigher"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/data_factory_linked_service_cosmosdb_mongoapi#timeouts DataFactoryLinkedServiceCosmosdbMongoapi#timeouts}
	Timeouts *DataFactoryLinkedServiceCosmosdbMongoapiTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

