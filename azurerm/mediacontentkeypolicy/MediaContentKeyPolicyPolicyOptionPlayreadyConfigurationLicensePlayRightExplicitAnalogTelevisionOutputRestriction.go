package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/media_content_key_policy#control_bits MediaContentKeyPolicy#control_bits}.
	ControlBits *float64 `field:"required" json:"controlBits" yaml:"controlBits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/media_content_key_policy#best_effort_enforced MediaContentKeyPolicy#best_effort_enforced}.
	BestEffortEnforced interface{} `field:"optional" json:"bestEffortEnforced" yaml:"bestEffortEnforced"`
}

