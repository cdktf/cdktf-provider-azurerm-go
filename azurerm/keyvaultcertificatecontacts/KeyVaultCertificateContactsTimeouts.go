package keyvaultcertificatecontacts


type KeyVaultCertificateContactsTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate_contacts#create KeyVaultCertificateContacts#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate_contacts#delete KeyVaultCertificateContacts#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate_contacts#read KeyVaultCertificateContacts#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate_contacts#update KeyVaultCertificateContacts#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
