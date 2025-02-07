// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazure


type DataFactoryIntegrationRuntimeAzureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_integration_runtime_azure#create DataFactoryIntegrationRuntimeAzure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_integration_runtime_azure#delete DataFactoryIntegrationRuntimeAzure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_integration_runtime_azure#read DataFactoryIntegrationRuntimeAzure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_integration_runtime_azure#update DataFactoryIntegrationRuntimeAzure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

