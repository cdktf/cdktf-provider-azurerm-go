package cdnendpoint


type CdnEndpointGlobalDeliveryRuleUrlRewriteAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#destination CdnEndpoint#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#source_pattern CdnEndpoint#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#preserve_unmatched_path CdnEndpoint#preserve_unmatched_path}.
	PreserveUnmatchedPath interface{} `field:"optional" json:"preserveUnmatchedPath" yaml:"preserveUnmatchedPath"`
}

