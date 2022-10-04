package automationsourcecontrol


type AutomationSourceControlSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_source_control#token AutomationSourceControl#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_source_control#token_type AutomationSourceControl#token_type}.
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_source_control#refresh_token AutomationSourceControl#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

