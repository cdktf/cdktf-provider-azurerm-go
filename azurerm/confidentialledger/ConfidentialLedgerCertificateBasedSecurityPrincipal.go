package confidentialledger


type ConfidentialLedgerCertificateBasedSecurityPrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/confidential_ledger#ledger_role_name ConfidentialLedger#ledger_role_name}.
	LedgerRoleName *string `field:"required" json:"ledgerRoleName" yaml:"ledgerRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/confidential_ledger#pem_public_key ConfidentialLedger#pem_public_key}.
	PemPublicKey *string `field:"required" json:"pemPublicKey" yaml:"pemPublicKey"`
}

