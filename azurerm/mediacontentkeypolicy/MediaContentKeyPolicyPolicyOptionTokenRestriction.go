package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionTokenRestriction struct {
	// alternate_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#alternate_key MediaContentKeyPolicy#alternate_key}
	AlternateKey interface{} `field:"optional" json:"alternateKey" yaml:"alternateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#audience MediaContentKeyPolicy#audience}.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#issuer MediaContentKeyPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#open_id_connect_discovery_document MediaContentKeyPolicy#open_id_connect_discovery_document}.
	OpenIdConnectDiscoveryDocument *string `field:"optional" json:"openIdConnectDiscoveryDocument" yaml:"openIdConnectDiscoveryDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#primary_rsa_token_key_exponent MediaContentKeyPolicy#primary_rsa_token_key_exponent}.
	PrimaryRsaTokenKeyExponent *string `field:"optional" json:"primaryRsaTokenKeyExponent" yaml:"primaryRsaTokenKeyExponent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#primary_rsa_token_key_modulus MediaContentKeyPolicy#primary_rsa_token_key_modulus}.
	PrimaryRsaTokenKeyModulus *string `field:"optional" json:"primaryRsaTokenKeyModulus" yaml:"primaryRsaTokenKeyModulus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#primary_symmetric_token_key MediaContentKeyPolicy#primary_symmetric_token_key}.
	PrimarySymmetricTokenKey *string `field:"optional" json:"primarySymmetricTokenKey" yaml:"primarySymmetricTokenKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#primary_x509_token_key_raw MediaContentKeyPolicy#primary_x509_token_key_raw}.
	PrimaryX509TokenKeyRaw *string `field:"optional" json:"primaryX509TokenKeyRaw" yaml:"primaryX509TokenKeyRaw"`
	// required_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#required_claim MediaContentKeyPolicy#required_claim}
	RequiredClaim interface{} `field:"optional" json:"requiredClaim" yaml:"requiredClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_content_key_policy#token_type MediaContentKeyPolicy#token_type}.
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

