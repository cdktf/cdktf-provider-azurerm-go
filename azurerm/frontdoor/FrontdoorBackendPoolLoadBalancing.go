package frontdoor


type FrontdoorBackendPoolLoadBalancing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#additional_latency_milliseconds Frontdoor#additional_latency_milliseconds}.
	AdditionalLatencyMilliseconds *float64 `field:"optional" json:"additionalLatencyMilliseconds" yaml:"additionalLatencyMilliseconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#sample_size Frontdoor#sample_size}.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#successful_samples_required Frontdoor#successful_samples_required}.
	SuccessfulSamplesRequired *float64 `field:"optional" json:"successfulSamplesRequired" yaml:"successfulSamplesRequired"`
}

