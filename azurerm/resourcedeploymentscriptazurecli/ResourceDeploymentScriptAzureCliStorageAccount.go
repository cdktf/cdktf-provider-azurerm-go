package resourcedeploymentscriptazurecli


type ResourceDeploymentScriptAzureCliStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#key ResourceDeploymentScriptAzureCli#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#name ResourceDeploymentScriptAzureCli#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
