package datafactorylinkedserviceazurefilestorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceAzureFileStorageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#connection_string DataFactoryLinkedServiceAzureFileStorage#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#data_factory_id DataFactoryLinkedServiceAzureFileStorage#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#name DataFactoryLinkedServiceAzureFileStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#additional_properties DataFactoryLinkedServiceAzureFileStorage#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#annotations DataFactoryLinkedServiceAzureFileStorage#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#description DataFactoryLinkedServiceAzureFileStorage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#file_share DataFactoryLinkedServiceAzureFileStorage#file_share}.
	FileShare *string `field:"optional" json:"fileShare" yaml:"fileShare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#host DataFactoryLinkedServiceAzureFileStorage#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#id DataFactoryLinkedServiceAzureFileStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#integration_runtime_name DataFactoryLinkedServiceAzureFileStorage#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#key_vault_password DataFactoryLinkedServiceAzureFileStorage#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceAzureFileStorageKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#parameters DataFactoryLinkedServiceAzureFileStorage#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#password DataFactoryLinkedServiceAzureFileStorage#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#timeouts DataFactoryLinkedServiceAzureFileStorage#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureFileStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_linked_service_azure_file_storage#user_id DataFactoryLinkedServiceAzureFileStorage#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

