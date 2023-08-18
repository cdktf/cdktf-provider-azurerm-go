package keyvaultcertificateissuer


type KeyVaultCertificateIssuerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/key_vault_certificate_issuer#create KeyVaultCertificateIssuer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/key_vault_certificate_issuer#delete KeyVaultCertificateIssuer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/key_vault_certificate_issuer#read KeyVaultCertificateIssuer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/key_vault_certificate_issuer#update KeyVaultCertificateIssuer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

