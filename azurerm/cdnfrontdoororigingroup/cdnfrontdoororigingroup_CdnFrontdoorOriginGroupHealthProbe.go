package cdnfrontdoororigingroup


type CdnFrontdoorOriginGroupHealthProbe struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#interval_in_seconds CdnFrontdoorOriginGroup#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"required" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#protocol CdnFrontdoorOriginGroup#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#path CdnFrontdoorOriginGroup#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#request_type CdnFrontdoorOriginGroup#request_type}.
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
}
