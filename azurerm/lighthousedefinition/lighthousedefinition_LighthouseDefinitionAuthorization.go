package lighthousedefinition


type LighthouseDefinitionAuthorization struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#principal_id LighthouseDefinition#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#role_definition_id LighthouseDefinition#role_definition_id}.
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#delegated_role_definition_ids LighthouseDefinition#delegated_role_definition_ids}.
	DelegatedRoleDefinitionIds *[]*string `field:"optional" json:"delegatedRoleDefinitionIds" yaml:"delegatedRoleDefinitionIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#principal_display_name LighthouseDefinition#principal_display_name}.
	PrincipalDisplayName *string `field:"optional" json:"principalDisplayName" yaml:"principalDisplayName"`
}
