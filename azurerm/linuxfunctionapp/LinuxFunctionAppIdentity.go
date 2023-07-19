package linuxfunctionapp


type LinuxFunctionAppIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_function_app#type LinuxFunctionApp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/linux_function_app#identity_ids LinuxFunctionApp#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

