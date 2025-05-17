// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazurefunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceAzureFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#data_factory_id DataFactoryLinkedServiceAzureFunction#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#name DataFactoryLinkedServiceAzureFunction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#url DataFactoryLinkedServiceAzureFunction#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#additional_properties DataFactoryLinkedServiceAzureFunction#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#annotations DataFactoryLinkedServiceAzureFunction#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#description DataFactoryLinkedServiceAzureFunction#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#id DataFactoryLinkedServiceAzureFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#integration_runtime_name DataFactoryLinkedServiceAzureFunction#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#key DataFactoryLinkedServiceAzureFunction#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// key_vault_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#key_vault_key DataFactoryLinkedServiceAzureFunction#key_vault_key}
	KeyVaultKey *DataFactoryLinkedServiceAzureFunctionKeyVaultKey `field:"optional" json:"keyVaultKey" yaml:"keyVaultKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#parameters DataFactoryLinkedServiceAzureFunction#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/data_factory_linked_service_azure_function#timeouts DataFactoryLinkedServiceAzureFunction#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureFunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

