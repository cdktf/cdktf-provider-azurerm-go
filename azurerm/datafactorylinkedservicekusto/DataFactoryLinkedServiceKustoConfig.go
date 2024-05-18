// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicekusto

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryLinkedServiceKustoConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#data_factory_id DataFactoryLinkedServiceKusto#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#kusto_database_name DataFactoryLinkedServiceKusto#kusto_database_name}.
	KustoDatabaseName *string `field:"required" json:"kustoDatabaseName" yaml:"kustoDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#kusto_endpoint DataFactoryLinkedServiceKusto#kusto_endpoint}.
	KustoEndpoint *string `field:"required" json:"kustoEndpoint" yaml:"kustoEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#name DataFactoryLinkedServiceKusto#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#additional_properties DataFactoryLinkedServiceKusto#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#annotations DataFactoryLinkedServiceKusto#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#description DataFactoryLinkedServiceKusto#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#id DataFactoryLinkedServiceKusto#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#integration_runtime_name DataFactoryLinkedServiceKusto#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#parameters DataFactoryLinkedServiceKusto#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#service_principal_id DataFactoryLinkedServiceKusto#service_principal_id}.
	ServicePrincipalId *string `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#service_principal_key DataFactoryLinkedServiceKusto#service_principal_key}.
	ServicePrincipalKey *string `field:"optional" json:"servicePrincipalKey" yaml:"servicePrincipalKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#tenant DataFactoryLinkedServiceKusto#tenant}.
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#timeouts DataFactoryLinkedServiceKusto#timeouts}
	Timeouts *DataFactoryLinkedServiceKustoTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_linked_service_kusto#use_managed_identity DataFactoryLinkedServiceKusto#use_managed_identity}.
	UseManagedIdentity interface{} `field:"optional" json:"useManagedIdentity" yaml:"useManagedIdentity"`
}

