package marketplaceagreement


type MarketplaceAgreementTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/marketplace_agreement#create MarketplaceAgreement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/marketplace_agreement#delete MarketplaceAgreement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/marketplace_agreement#read MarketplaceAgreement#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/marketplace_agreement#update MarketplaceAgreement#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
