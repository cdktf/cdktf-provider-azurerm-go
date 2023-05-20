package blueprintassignment


type BlueprintAssignmentIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/blueprint_assignment#identity_ids BlueprintAssignment#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/blueprint_assignment#type BlueprintAssignment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

