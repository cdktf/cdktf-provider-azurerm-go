// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#agc_and_color_stripe_restriction MediaContentKeyPolicy#agc_and_color_stripe_restriction}.
	AgcAndColorStripeRestriction *float64 `field:"optional" json:"agcAndColorStripeRestriction" yaml:"agcAndColorStripeRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#allow_passing_video_content_to_unknown_output MediaContentKeyPolicy#allow_passing_video_content_to_unknown_output}.
	AllowPassingVideoContentToUnknownOutput *string `field:"optional" json:"allowPassingVideoContentToUnknownOutput" yaml:"allowPassingVideoContentToUnknownOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#analog_video_opl MediaContentKeyPolicy#analog_video_opl}.
	AnalogVideoOpl *float64 `field:"optional" json:"analogVideoOpl" yaml:"analogVideoOpl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#compressed_digital_audio_opl MediaContentKeyPolicy#compressed_digital_audio_opl}.
	CompressedDigitalAudioOpl *float64 `field:"optional" json:"compressedDigitalAudioOpl" yaml:"compressedDigitalAudioOpl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#compressed_digital_video_opl MediaContentKeyPolicy#compressed_digital_video_opl}.
	CompressedDigitalVideoOpl *float64 `field:"optional" json:"compressedDigitalVideoOpl" yaml:"compressedDigitalVideoOpl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#digital_video_only_content_restriction MediaContentKeyPolicy#digital_video_only_content_restriction}.
	DigitalVideoOnlyContentRestriction interface{} `field:"optional" json:"digitalVideoOnlyContentRestriction" yaml:"digitalVideoOnlyContentRestriction"`
	// explicit_analog_television_output_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#explicit_analog_television_output_restriction MediaContentKeyPolicy#explicit_analog_television_output_restriction}
	ExplicitAnalogTelevisionOutputRestriction *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction `field:"optional" json:"explicitAnalogTelevisionOutputRestriction" yaml:"explicitAnalogTelevisionOutputRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#first_play_expiration MediaContentKeyPolicy#first_play_expiration}.
	FirstPlayExpiration *string `field:"optional" json:"firstPlayExpiration" yaml:"firstPlayExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#image_constraint_for_analog_component_video_restriction MediaContentKeyPolicy#image_constraint_for_analog_component_video_restriction}.
	ImageConstraintForAnalogComponentVideoRestriction interface{} `field:"optional" json:"imageConstraintForAnalogComponentVideoRestriction" yaml:"imageConstraintForAnalogComponentVideoRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#image_constraint_for_analog_computer_monitor_restriction MediaContentKeyPolicy#image_constraint_for_analog_computer_monitor_restriction}.
	ImageConstraintForAnalogComputerMonitorRestriction interface{} `field:"optional" json:"imageConstraintForAnalogComputerMonitorRestriction" yaml:"imageConstraintForAnalogComputerMonitorRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#scms_restriction MediaContentKeyPolicy#scms_restriction}.
	ScmsRestriction *float64 `field:"optional" json:"scmsRestriction" yaml:"scmsRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#uncompressed_digital_audio_opl MediaContentKeyPolicy#uncompressed_digital_audio_opl}.
	UncompressedDigitalAudioOpl *float64 `field:"optional" json:"uncompressedDigitalAudioOpl" yaml:"uncompressedDigitalAudioOpl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/media_content_key_policy#uncompressed_digital_video_opl MediaContentKeyPolicy#uncompressed_digital_video_opl}.
	UncompressedDigitalVideoOpl *float64 `field:"optional" json:"uncompressedDigitalVideoOpl" yaml:"uncompressedDigitalVideoOpl"`
}

