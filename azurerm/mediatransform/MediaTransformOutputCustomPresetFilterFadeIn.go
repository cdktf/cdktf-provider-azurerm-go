package mediatransform


type MediaTransformOutputCustomPresetFilterFadeIn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/media_transform#duration MediaTransform#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/media_transform#fade_color MediaTransform#fade_color}.
	FadeColor *string `field:"required" json:"fadeColor" yaml:"fadeColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/media_transform#start MediaTransform#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

