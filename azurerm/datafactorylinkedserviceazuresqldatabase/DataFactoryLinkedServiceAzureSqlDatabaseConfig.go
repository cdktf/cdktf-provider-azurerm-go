// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuresqldatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceAzureSqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#data_factory_id DataFactoryLinkedServiceAzureSqlDatabase#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#name DataFactoryLinkedServiceAzureSqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#additional_properties DataFactoryLinkedServiceAzureSqlDatabase#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#annotations DataFactoryLinkedServiceAzureSqlDatabase#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#connection_string DataFactoryLinkedServiceAzureSqlDatabase#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#credential_name DataFactoryLinkedServiceAzureSqlDatabase#credential_name}.
	CredentialName *string `field:"optional" json:"credentialName" yaml:"credentialName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#description DataFactoryLinkedServiceAzureSqlDatabase#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#id DataFactoryLinkedServiceAzureSqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#integration_runtime_name DataFactoryLinkedServiceAzureSqlDatabase#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#key_vault_connection_string DataFactoryLinkedServiceAzureSqlDatabase#key_vault_connection_string}
	KeyVaultConnectionString *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString `field:"optional" json:"keyVaultConnectionString" yaml:"keyVaultConnectionString"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#key_vault_password DataFactoryLinkedServiceAzureSqlDatabase#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#parameters DataFactoryLinkedServiceAzureSqlDatabase#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#service_principal_id DataFactoryLinkedServiceAzureSqlDatabase#service_principal_id}.
	ServicePrincipalId *string `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#service_principal_key DataFactoryLinkedServiceAzureSqlDatabase#service_principal_key}.
	ServicePrincipalKey *string `field:"optional" json:"servicePrincipalKey" yaml:"servicePrincipalKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#tenant_id DataFactoryLinkedServiceAzureSqlDatabase#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#timeouts DataFactoryLinkedServiceAzureSqlDatabase#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureSqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/data_factory_linked_service_azure_sql_database#use_managed_identity DataFactoryLinkedServiceAzureSqlDatabase#use_managed_identity}.
	UseManagedIdentity interface{} `field:"optional" json:"useManagedIdentity" yaml:"useManagedIdentity"`
}

