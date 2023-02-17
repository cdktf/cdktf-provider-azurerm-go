package confidentialledger


type ConfidentialLedgerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/confidential_ledger#create ConfidentialLedger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/confidential_ledger#delete ConfidentialLedger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/confidential_ledger#read ConfidentialLedger#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/confidential_ledger#update ConfidentialLedger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

