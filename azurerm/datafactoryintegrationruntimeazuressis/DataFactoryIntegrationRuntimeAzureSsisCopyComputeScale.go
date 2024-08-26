// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory_integration_runtime_azure_ssis#data_integration_unit DataFactoryIntegrationRuntimeAzureSsis#data_integration_unit}.
	DataIntegrationUnit *float64 `field:"optional" json:"dataIntegrationUnit" yaml:"dataIntegrationUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory_integration_runtime_azure_ssis#time_to_live DataFactoryIntegrationRuntimeAzureSsis#time_to_live}.
	TimeToLive *float64 `field:"optional" json:"timeToLive" yaml:"timeToLive"`
}

