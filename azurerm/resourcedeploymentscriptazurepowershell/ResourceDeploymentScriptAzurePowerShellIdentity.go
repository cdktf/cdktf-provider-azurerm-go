package resourcedeploymentscriptazurepowershell


type ResourceDeploymentScriptAzurePowerShellIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/resource_deployment_script_azure_power_shell#identity_ids ResourceDeploymentScriptAzurePowerShell#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/resource_deployment_script_azure_power_shell#type ResourceDeploymentScriptAzurePowerShell#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

