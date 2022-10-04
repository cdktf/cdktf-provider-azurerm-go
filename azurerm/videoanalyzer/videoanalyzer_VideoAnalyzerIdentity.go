package videoanalyzer


type VideoAnalyzerIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/video_analyzer#identity_ids VideoAnalyzer#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/video_analyzer#type VideoAnalyzer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

