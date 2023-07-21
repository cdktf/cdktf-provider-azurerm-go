package mediajob


type MediaJobInputAsset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/media_job#name MediaJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/media_job#label MediaJob#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

