package mediatransform


type MediaTransformOutputCustomPresetCodecPngImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#start MediaTransform#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#key_frame_interval MediaTransform#key_frame_interval}.
	KeyFrameInterval *string `field:"optional" json:"keyFrameInterval" yaml:"keyFrameInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#label MediaTransform#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// layer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#layer MediaTransform#layer}
	Layer interface{} `field:"optional" json:"layer" yaml:"layer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#range MediaTransform#range}.
	Range *string `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#step MediaTransform#step}.
	Step *string `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#stretch_mode MediaTransform#stretch_mode}.
	StretchMode *string `field:"optional" json:"stretchMode" yaml:"stretchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#sync_mode MediaTransform#sync_mode}.
	SyncMode *string `field:"optional" json:"syncMode" yaml:"syncMode"`
}

