package cdnfrontdoorrule


type CdnFrontdoorRuleActionsUrlRewriteAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#destination CdnFrontdoorRule#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#source_pattern CdnFrontdoorRule#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#preserve_unmatched_path CdnFrontdoorRule#preserve_unmatched_path}.
	PreserveUnmatchedPath interface{} `field:"optional" json:"preserveUnmatchedPath" yaml:"preserveUnmatchedPath"`
}

