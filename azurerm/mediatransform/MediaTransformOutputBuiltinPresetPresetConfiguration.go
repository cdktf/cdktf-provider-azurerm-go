package mediatransform


type MediaTransformOutputBuiltinPresetPresetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#complexity MediaTransform#complexity}.
	Complexity *string `field:"optional" json:"complexity" yaml:"complexity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#interleave_output MediaTransform#interleave_output}.
	InterleaveOutput *string `field:"optional" json:"interleaveOutput" yaml:"interleaveOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#key_frame_interval_in_seconds MediaTransform#key_frame_interval_in_seconds}.
	KeyFrameIntervalInSeconds *float64 `field:"optional" json:"keyFrameIntervalInSeconds" yaml:"keyFrameIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#max_bitrate_bps MediaTransform#max_bitrate_bps}.
	MaxBitrateBps *float64 `field:"optional" json:"maxBitrateBps" yaml:"maxBitrateBps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#max_height MediaTransform#max_height}.
	MaxHeight *float64 `field:"optional" json:"maxHeight" yaml:"maxHeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#max_layers MediaTransform#max_layers}.
	MaxLayers *float64 `field:"optional" json:"maxLayers" yaml:"maxLayers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#min_bitrate_bps MediaTransform#min_bitrate_bps}.
	MinBitrateBps *float64 `field:"optional" json:"minBitrateBps" yaml:"minBitrateBps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/media_transform#min_height MediaTransform#min_height}.
	MinHeight *float64 `field:"optional" json:"minHeight" yaml:"minHeight"`
}

