package sharedimage


type SharedImagePurchasePlan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#name SharedImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#product SharedImage#product}.
	Product *string `field:"optional" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#publisher SharedImage#publisher}.
	Publisher *string `field:"optional" json:"publisher" yaml:"publisher"`
}

