package mediatransform


type MediaTransformOutputCustomPreset struct {
	// codec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/media_transform#codec MediaTransform#codec}
	Codec interface{} `field:"required" json:"codec" yaml:"codec"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/media_transform#format MediaTransform#format}
	Format interface{} `field:"required" json:"format" yaml:"format"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/media_transform#filter MediaTransform#filter}
	Filter *MediaTransformOutputCustomPresetFilter `field:"optional" json:"filter" yaml:"filter"`
}

