package cdnfrontdoororigingroup


type CdnFrontdoorOriginGroupLoadBalancing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#additional_latency_in_milliseconds CdnFrontdoorOriginGroup#additional_latency_in_milliseconds}.
	AdditionalLatencyInMilliseconds *float64 `field:"optional" json:"additionalLatencyInMilliseconds" yaml:"additionalLatencyInMilliseconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#sample_size CdnFrontdoorOriginGroup#sample_size}.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_origin_group#successful_samples_required CdnFrontdoorOriginGroup#successful_samples_required}.
	SuccessfulSamplesRequired *float64 `field:"optional" json:"successfulSamplesRequired" yaml:"successfulSamplesRequired"`
}
