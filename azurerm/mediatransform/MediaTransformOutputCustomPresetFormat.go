package mediatransform


type MediaTransformOutputCustomPresetFormat struct {
	// mp4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/media_transform#mp4 MediaTransform#mp4}
	Mp4 *MediaTransformOutputCustomPresetFormatMp4 `field:"optional" json:"mp4" yaml:"mp4"`
	// transport_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/media_transform#transport_stream MediaTransform#transport_stream}
	TransportStream *MediaTransformOutputCustomPresetFormatTransportStream `field:"optional" json:"transportStream" yaml:"transportStream"`
}

