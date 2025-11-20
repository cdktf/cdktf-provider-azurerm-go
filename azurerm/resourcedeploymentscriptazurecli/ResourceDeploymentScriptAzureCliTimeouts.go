// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcedeploymentscriptazurecli


type ResourceDeploymentScriptAzureCliTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/resource_deployment_script_azure_cli#create ResourceDeploymentScriptAzureCli#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/resource_deployment_script_azure_cli#delete ResourceDeploymentScriptAzureCli#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/resource_deployment_script_azure_cli#read ResourceDeploymentScriptAzureCli#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/resource_deployment_script_azure_cli#update ResourceDeploymentScriptAzureCli#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

