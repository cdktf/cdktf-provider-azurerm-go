package resourcedeploymentscriptazurecli


type ResourceDeploymentScriptAzureCliIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#identity_ids ResourceDeploymentScriptAzureCli#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_deployment_script_azure_cli#type ResourceDeploymentScriptAzureCli#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}
