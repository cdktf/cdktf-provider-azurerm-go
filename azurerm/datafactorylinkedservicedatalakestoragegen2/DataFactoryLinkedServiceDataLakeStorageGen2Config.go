// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicedatalakestoragegen2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceDataLakeStorageGen2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#data_factory_id DataFactoryLinkedServiceDataLakeStorageGen2#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#name DataFactoryLinkedServiceDataLakeStorageGen2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#url DataFactoryLinkedServiceDataLakeStorageGen2#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#additional_properties DataFactoryLinkedServiceDataLakeStorageGen2#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#annotations DataFactoryLinkedServiceDataLakeStorageGen2#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#description DataFactoryLinkedServiceDataLakeStorageGen2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#id DataFactoryLinkedServiceDataLakeStorageGen2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#integration_runtime_name DataFactoryLinkedServiceDataLakeStorageGen2#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#parameters DataFactoryLinkedServiceDataLakeStorageGen2#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#service_principal_id DataFactoryLinkedServiceDataLakeStorageGen2#service_principal_id}.
	ServicePrincipalId *string `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#service_principal_key DataFactoryLinkedServiceDataLakeStorageGen2#service_principal_key}.
	ServicePrincipalKey *string `field:"optional" json:"servicePrincipalKey" yaml:"servicePrincipalKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#storage_account_key DataFactoryLinkedServiceDataLakeStorageGen2#storage_account_key}.
	StorageAccountKey *string `field:"optional" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#tenant DataFactoryLinkedServiceDataLakeStorageGen2#tenant}.
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#timeouts DataFactoryLinkedServiceDataLakeStorageGen2#timeouts}
	Timeouts *DataFactoryLinkedServiceDataLakeStorageGen2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_linked_service_data_lake_storage_gen2#use_managed_identity DataFactoryLinkedServiceDataLakeStorageGen2#use_managed_identity}.
	UseManagedIdentity interface{} `field:"optional" json:"useManagedIdentity" yaml:"useManagedIdentity"`
}

