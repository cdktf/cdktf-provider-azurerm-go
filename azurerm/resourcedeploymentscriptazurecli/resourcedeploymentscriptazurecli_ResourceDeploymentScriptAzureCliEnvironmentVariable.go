package resourcedeploymentscriptazurecli


type ResourceDeploymentScriptAzureCliEnvironmentVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#name ResourceDeploymentScriptAzureCli#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#secure_value ResourceDeploymentScriptAzureCli#secure_value}.
	SecureValue *string `field:"optional" json:"secureValue" yaml:"secureValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#value ResourceDeploymentScriptAzureCli#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

