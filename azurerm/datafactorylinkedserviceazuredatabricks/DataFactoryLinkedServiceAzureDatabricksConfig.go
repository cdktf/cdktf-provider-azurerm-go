// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuredatabricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceAzureDatabricksConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#adb_domain DataFactoryLinkedServiceAzureDatabricks#adb_domain}.
	AdbDomain *string `field:"required" json:"adbDomain" yaml:"adbDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#data_factory_id DataFactoryLinkedServiceAzureDatabricks#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#name DataFactoryLinkedServiceAzureDatabricks#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#access_token DataFactoryLinkedServiceAzureDatabricks#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#additional_properties DataFactoryLinkedServiceAzureDatabricks#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#annotations DataFactoryLinkedServiceAzureDatabricks#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#description DataFactoryLinkedServiceAzureDatabricks#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#existing_cluster_id DataFactoryLinkedServiceAzureDatabricks#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#id DataFactoryLinkedServiceAzureDatabricks#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#instance_pool DataFactoryLinkedServiceAzureDatabricks#instance_pool}
	InstancePool *DataFactoryLinkedServiceAzureDatabricksInstancePool `field:"optional" json:"instancePool" yaml:"instancePool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#integration_runtime_name DataFactoryLinkedServiceAzureDatabricks#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#key_vault_password DataFactoryLinkedServiceAzureDatabricks#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#msi_work_space_resource_id DataFactoryLinkedServiceAzureDatabricks#msi_work_space_resource_id}.
	MsiWorkSpaceResourceId *string `field:"optional" json:"msiWorkSpaceResourceId" yaml:"msiWorkSpaceResourceId"`
	// new_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#new_cluster_config DataFactoryLinkedServiceAzureDatabricks#new_cluster_config}
	NewClusterConfig *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig `field:"optional" json:"newClusterConfig" yaml:"newClusterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#parameters DataFactoryLinkedServiceAzureDatabricks#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/data_factory_linked_service_azure_databricks#timeouts DataFactoryLinkedServiceAzureDatabricks#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureDatabricksTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

