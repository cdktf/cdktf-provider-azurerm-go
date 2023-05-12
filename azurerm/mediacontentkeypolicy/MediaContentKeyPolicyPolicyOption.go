package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#name MediaContentKeyPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#clear_key_configuration_enabled MediaContentKeyPolicy#clear_key_configuration_enabled}.
	ClearKeyConfigurationEnabled interface{} `field:"optional" json:"clearKeyConfigurationEnabled" yaml:"clearKeyConfigurationEnabled"`
	// fairplay_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#fairplay_configuration MediaContentKeyPolicy#fairplay_configuration}
	FairplayConfiguration *MediaContentKeyPolicyPolicyOptionFairplayConfiguration `field:"optional" json:"fairplayConfiguration" yaml:"fairplayConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#open_restriction_enabled MediaContentKeyPolicy#open_restriction_enabled}.
	OpenRestrictionEnabled interface{} `field:"optional" json:"openRestrictionEnabled" yaml:"openRestrictionEnabled"`
	// playready_configuration_license block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#playready_configuration_license MediaContentKeyPolicy#playready_configuration_license}
	PlayreadyConfigurationLicense interface{} `field:"optional" json:"playreadyConfigurationLicense" yaml:"playreadyConfigurationLicense"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#playready_response_custom_data MediaContentKeyPolicy#playready_response_custom_data}.
	PlayreadyResponseCustomData *string `field:"optional" json:"playreadyResponseCustomData" yaml:"playreadyResponseCustomData"`
	// token_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#token_restriction MediaContentKeyPolicy#token_restriction}
	TokenRestriction *MediaContentKeyPolicyPolicyOptionTokenRestriction `field:"optional" json:"tokenRestriction" yaml:"tokenRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/media_content_key_policy#widevine_configuration_template MediaContentKeyPolicy#widevine_configuration_template}.
	WidevineConfigurationTemplate *string `field:"optional" json:"widevineConfigurationTemplate" yaml:"widevineConfigurationTemplate"`
}

