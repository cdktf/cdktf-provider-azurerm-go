package mediatransform


type MediaTransformOutputCustomPresetCodecH264VideoLayer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#bitrate MediaTransform#bitrate}.
	Bitrate *float64 `field:"required" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#adaptive_b_frame_enabled MediaTransform#adaptive_b_frame_enabled}.
	AdaptiveBFrameEnabled interface{} `field:"optional" json:"adaptiveBFrameEnabled" yaml:"adaptiveBFrameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#b_frames MediaTransform#b_frames}.
	BFrames *float64 `field:"optional" json:"bFrames" yaml:"bFrames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#buffer_window MediaTransform#buffer_window}.
	BufferWindow *string `field:"optional" json:"bufferWindow" yaml:"bufferWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#crf MediaTransform#crf}.
	Crf *float64 `field:"optional" json:"crf" yaml:"crf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#entropy_mode MediaTransform#entropy_mode}.
	EntropyMode *string `field:"optional" json:"entropyMode" yaml:"entropyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#frame_rate MediaTransform#frame_rate}.
	FrameRate *string `field:"optional" json:"frameRate" yaml:"frameRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#height MediaTransform#height}.
	Height *string `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#label MediaTransform#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#level MediaTransform#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#max_bitrate MediaTransform#max_bitrate}.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#profile MediaTransform#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#reference_frames MediaTransform#reference_frames}.
	ReferenceFrames *float64 `field:"optional" json:"referenceFrames" yaml:"referenceFrames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#slices MediaTransform#slices}.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/media_transform#width MediaTransform#width}.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

