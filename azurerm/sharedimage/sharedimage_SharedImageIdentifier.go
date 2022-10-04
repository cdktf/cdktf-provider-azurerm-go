package sharedimage


type SharedImageIdentifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#offer SharedImage#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#publisher SharedImage#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#sku SharedImage#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}

