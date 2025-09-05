// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcedeploymentscriptazurecli


type ResourceDeploymentScriptAzureCliStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/resource_deployment_script_azure_cli#key ResourceDeploymentScriptAzureCli#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/resource_deployment_script_azure_cli#name ResourceDeploymentScriptAzureCli#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

