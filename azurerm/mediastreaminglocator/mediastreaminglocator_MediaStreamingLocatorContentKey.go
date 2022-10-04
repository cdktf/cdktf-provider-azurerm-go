package mediastreaminglocator


type MediaStreamingLocatorContentKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator#content_key_id MediaStreamingLocator#content_key_id}.
	ContentKeyId *string `field:"optional" json:"contentKeyId" yaml:"contentKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator#label_reference_in_streaming_policy MediaStreamingLocator#label_reference_in_streaming_policy}.
	LabelReferenceInStreamingPolicy *string `field:"optional" json:"labelReferenceInStreamingPolicy" yaml:"labelReferenceInStreamingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator#policy_name MediaStreamingLocator#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator#type MediaStreamingLocator#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_locator#value MediaStreamingLocator#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

