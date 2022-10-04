package keyvaultcertificate


type KeyVaultCertificateCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#contents KeyVaultCertificate#contents}.
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#password KeyVaultCertificate#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
}

