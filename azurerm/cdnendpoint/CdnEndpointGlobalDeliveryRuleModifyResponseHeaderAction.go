package cdnendpoint


type CdnEndpointGlobalDeliveryRuleModifyResponseHeaderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#value CdnEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

