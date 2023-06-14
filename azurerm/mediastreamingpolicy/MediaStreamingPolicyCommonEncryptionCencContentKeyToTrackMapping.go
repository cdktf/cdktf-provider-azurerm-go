package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCencContentKeyToTrackMapping struct {
	// track block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_streaming_policy#track MediaStreamingPolicy#track}
	Track interface{} `field:"required" json:"track" yaml:"track"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_streaming_policy#label MediaStreamingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/media_streaming_policy#policy_name MediaStreamingPolicy#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

