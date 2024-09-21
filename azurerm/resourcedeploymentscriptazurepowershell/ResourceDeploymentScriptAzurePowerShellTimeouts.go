// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcedeploymentscriptazurepowershell


type ResourceDeploymentScriptAzurePowerShellTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/resource_deployment_script_azure_power_shell#create ResourceDeploymentScriptAzurePowerShell#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/resource_deployment_script_azure_power_shell#delete ResourceDeploymentScriptAzurePowerShell#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/resource_deployment_script_azure_power_shell#read ResourceDeploymentScriptAzurePowerShell#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/resource_deployment_script_azure_power_shell#update ResourceDeploymentScriptAzurePowerShell#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

