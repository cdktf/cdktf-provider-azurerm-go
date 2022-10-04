package mediajob


type MediaJobInputAsset struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_job#name MediaJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_job#label MediaJob#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

