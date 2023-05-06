package containergroup


type ContainerGroupImageRegistryCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_group#server ContainerGroup#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_group#password ContainerGroup#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The User Assigned Identity to use for Container Registry access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_group#user_assigned_identity_id ContainerGroup#user_assigned_identity_id}
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_group#username ContainerGroup#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

