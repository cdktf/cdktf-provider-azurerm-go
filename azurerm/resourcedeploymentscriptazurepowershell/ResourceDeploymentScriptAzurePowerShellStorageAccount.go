package resourcedeploymentscriptazurepowershell


type ResourceDeploymentScriptAzurePowerShellStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_power_shell#key ResourceDeploymentScriptAzurePowerShell#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_power_shell#name ResourceDeploymentScriptAzurePowerShell#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

