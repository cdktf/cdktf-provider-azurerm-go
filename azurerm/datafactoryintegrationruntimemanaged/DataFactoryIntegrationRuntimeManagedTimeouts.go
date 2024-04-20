// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimemanaged


type DataFactoryIntegrationRuntimeManagedTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_integration_runtime_managed#create DataFactoryIntegrationRuntimeManaged#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_integration_runtime_managed#delete DataFactoryIntegrationRuntimeManaged#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_integration_runtime_managed#read DataFactoryIntegrationRuntimeManaged#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_integration_runtime_managed#update DataFactoryIntegrationRuntimeManaged#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

