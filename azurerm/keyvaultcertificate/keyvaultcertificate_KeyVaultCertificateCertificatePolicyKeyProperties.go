package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyKeyProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#exportable KeyVaultCertificate#exportable}.
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#key_type KeyVaultCertificate#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#reuse_key KeyVaultCertificate#reuse_key}.
	ReuseKey interface{} `field:"required" json:"reuseKey" yaml:"reuseKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#curve KeyVaultCertificate#curve}.
	Curve *string `field:"optional" json:"curve" yaml:"curve"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_certificate#key_size KeyVaultCertificate#key_size}.
	KeySize *float64 `field:"optional" json:"keySize" yaml:"keySize"`
}
