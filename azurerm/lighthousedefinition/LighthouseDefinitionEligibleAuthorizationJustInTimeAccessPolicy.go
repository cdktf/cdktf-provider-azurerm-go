package lighthousedefinition


type LighthouseDefinitionEligibleAuthorizationJustInTimeAccessPolicy struct {
	// approver block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#approver LighthouseDefinition#approver}
	Approver interface{} `field:"optional" json:"approver" yaml:"approver"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#maximum_activation_duration LighthouseDefinition#maximum_activation_duration}.
	MaximumActivationDuration *string `field:"optional" json:"maximumActivationDuration" yaml:"maximumActivationDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#multi_factor_auth_provider LighthouseDefinition#multi_factor_auth_provider}.
	MultiFactorAuthProvider *string `field:"optional" json:"multiFactorAuthProvider" yaml:"multiFactorAuthProvider"`
}
