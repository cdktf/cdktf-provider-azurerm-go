package resourcedeploymentscriptazurepowershell


type ResourceDeploymentScriptAzurePowerShellEnvironmentVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_power_shell#name ResourceDeploymentScriptAzurePowerShell#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_power_shell#secure_value ResourceDeploymentScriptAzurePowerShell#secure_value}.
	SecureValue *string `field:"optional" json:"secureValue" yaml:"secureValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_power_shell#value ResourceDeploymentScriptAzurePowerShell#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}
