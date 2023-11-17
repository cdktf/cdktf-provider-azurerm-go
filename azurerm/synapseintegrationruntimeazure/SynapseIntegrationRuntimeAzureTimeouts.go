// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseintegrationruntimeazure


type SynapseIntegrationRuntimeAzureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/synapse_integration_runtime_azure#create SynapseIntegrationRuntimeAzure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/synapse_integration_runtime_azure#delete SynapseIntegrationRuntimeAzure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/synapse_integration_runtime_azure#read SynapseIntegrationRuntimeAzure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/synapse_integration_runtime_azure#update SynapseIntegrationRuntimeAzure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

