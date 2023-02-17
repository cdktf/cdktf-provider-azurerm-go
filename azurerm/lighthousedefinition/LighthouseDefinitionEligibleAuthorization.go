package lighthousedefinition


type LighthouseDefinitionEligibleAuthorization struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#principal_id LighthouseDefinition#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#role_definition_id LighthouseDefinition#role_definition_id}.
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// just_in_time_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#just_in_time_access_policy LighthouseDefinition#just_in_time_access_policy}
	JustInTimeAccessPolicy *LighthouseDefinitionEligibleAuthorizationJustInTimeAccessPolicy `field:"optional" json:"justInTimeAccessPolicy" yaml:"justInTimeAccessPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#principal_display_name LighthouseDefinition#principal_display_name}.
	PrincipalDisplayName *string `field:"optional" json:"principalDisplayName" yaml:"principalDisplayName"`
}

