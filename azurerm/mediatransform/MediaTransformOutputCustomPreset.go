package mediatransform


type MediaTransformOutputCustomPreset struct {
	// codec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#codec MediaTransform#codec}
	Codec interface{} `field:"required" json:"codec" yaml:"codec"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#format MediaTransform#format}
	Format interface{} `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#experimental_options MediaTransform#experimental_options}.
	ExperimentalOptions *map[string]*string `field:"optional" json:"experimentalOptions" yaml:"experimentalOptions"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#filter MediaTransform#filter}
	Filter *MediaTransformOutputCustomPresetFilter `field:"optional" json:"filter" yaml:"filter"`
}

